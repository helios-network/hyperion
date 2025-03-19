package solidity

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ctypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Helios-Chain-Labs/etherman/deployer"

	"github.com/Helios-Chain-Labs/etherman/sol"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/hyperion"
	wrappers "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
)

var _ = Describe("Contract Tests", func() {
	_ = Describe("Hyperion", func() {
		var (
			hyperionTxOpts   deployer.ContractTxOpts
			hyperionCallOpts deployer.ContractCallOpts
			hyperionLogsOpts deployer.ContractLogsOpts
			hyperionContract *sol.Contract

			deployArgs   deployer.AbiMethodInputMapperFunc
			deployErr    error
			deployTxHash common.Hash

			hyperionID common.Hash
			minPower   *big.Int
			validators []common.Address
			powers     []*big.Int
		)

		BeforeEach(func() {
			hyperionID = formatBytes32String("foo")
			validators = getEthAddresses(CosmosAccounts[:3]...)
			minPower = big.NewInt(3500)
			powers = []*big.Int{
				big.NewInt(3000),
				big.NewInt(1500),
				big.NewInt(500),
			}

			deployArgs = withArgsFn(
				hyperionID,
				minPower,
				validators,
				powers,
			)
		})

		JustBeforeEach(func() {
			// don't redeploy if already deployed
			if hyperionContract != nil {
				return
			}

			hyperionDeployOpts := deployer.ContractDeployOpts{
				From:          EthAccounts[0].EthAddress,
				FromPk:        EthAccounts[0].EthPrivKey,
				SolSource:     "../../../Ethereum-Bridge-Contract/contracts/Hyperion.sol",
				ContractName:  "Hyperion",
				Await:         true,
				CoverageAgent: CoverageAgent,
			}

			_, hyperionContract, deployErr = ContractDeployer.Deploy(context.Background(), hyperionDeployOpts, noArgs)
			orFail(deployErr)

			hyperionTxOpts = deployer.ContractTxOpts{
				From:          EthAccounts[0].EthAddress,
				FromPk:        EthAccounts[0].EthPrivKey,
				SolSource:     "../../../Ethereum-Bridge-Contract/contracts/Hyperion.sol",
				ContractName:  "Hyperion",
				Contract:      hyperionContract.Address,
				Await:         true,
				CoverageAgent: CoverageAgent,
			}

			deployTxHash, _, deployErr = ContractDeployer.Tx(context.Background(), hyperionTxOpts, "initialize", deployArgs)
		})

		_ = Context("Contract fails to initialize", func() {
			AfterEach(func() {
				deployArgs = withArgsFn(
					hyperionID,
					minPower,
					validators,
					powers,
				)

				// force redeployment
				hyperionContract = nil
			})

			_ = When("Throws on malformed valset", func() {
				BeforeEach(func() {
					deployArgs = withArgsFn(
						hyperionID,
						minPower,
						validators,
						powers[:1], // only one
					)
				})

				It("Should fail with error", func() {
					Ω(deployErr).ShouldNot(BeNil())
					Ω(deployErr.Error()).Should(ContainSubstring("Malformed current validator set"))
				})
			})

			_ = When("Throws on insufficient power", func() {
				BeforeEach(func() {
					deployArgs = withArgsFn(
						hyperionID,
						big.NewInt(10000),
						validators,
						powers,
					)
				})

				It("Should fail with error", func() {
					Ω(deployErr).ShouldNot(BeNil())
					Ω(deployErr.Error()).Should(ContainSubstring("Submitted validator set signatures do not have enough power"))
				})
			})
		})

		_ = Context("Hyperion contract deployment and initialization done", func() {
			var (
				hyperionOwner Account
			)

			BeforeEach(func() {
				hyperionOwner = EthAccounts[0]
			})

			JustBeforeEach(func() {
				orFail(deployErr)

				hyperionTxOpts = deployer.ContractTxOpts{
					From:          hyperionOwner.EthAddress,
					FromPk:        hyperionOwner.EthPrivKey,
					SolSource:     "../../../Ethereum-Bridge-Contract/contracts/Hyperion.sol",
					ContractName:  "Hyperion",
					Contract:      hyperionContract.Address,
					Await:         true,
					CoverageAgent: CoverageAgent,
				}

				hyperionCallOpts = deployer.ContractCallOpts{
					From:          hyperionOwner.EthAddress,
					SolSource:     "../../../Ethereum-Bridge-Contract/contracts/Hyperion.sol",
					ContractName:  "Hyperion",
					Contract:      hyperionContract.Address,
					CoverageAgent: CoverageAgent,
					CoverageCall: deployer.ContractCoverageCallOpts{
						FromPk: hyperionOwner.EthPrivKey,
					},
				}

				hyperionLogsOpts = deployer.ContractLogsOpts{
					SolSource:     "../../../Ethereum-Bridge-Contract/contracts/Hyperion.sol",
					ContractName:  "Hyperion",
					Contract:      hyperionContract.Address,
					CoverageAgent: CoverageAgent,
				}
			})

			_ = Describe("Check contract state", func() {
				It("Should have address", func() {
					Ω(hyperionTxOpts.Contract).ShouldNot(Equal(zeroAddress))
					Ω(hyperionCallOpts.Contract).ShouldNot(Equal(zeroAddress))
				})

				It("Should have valid power threshold", func() {
					var state_powerThreshold *big.Int

					out, outAbi, err := ContractDeployer.Call(context.Background(), hyperionCallOpts,
						"state_powerThreshold", noArgs,
					)
					Ω(err).Should(BeNil())

					err = outAbi.Copy(&state_powerThreshold, out)
					Ω(err).Should(BeNil())
					Ω(state_powerThreshold.String()).Should(Equal(minPower.String()))
				})

				It("Should have valid hyperionId", func() {
					var state_hyperionId common.Hash

					out, outAbi, err := ContractDeployer.Call(context.Background(), hyperionCallOpts,
						"state_hyperionId", noArgs,
					)
					Ω(err).Should(BeNil())

					err = outAbi.Copy(&state_hyperionId, out)
					Ω(err).Should(BeNil())
					Ω(state_hyperionId).Should(Equal(hyperionID))
				})

				It("Should have generated a valid checkpoint", func() {
					var state_lastValsetCheckpoint common.Hash

					out, outAbi, err := ContractDeployer.Call(context.Background(), hyperionCallOpts,
						"state_lastValsetCheckpoint", noArgs,
					)
					Ω(err).Should(BeNil())

					offchainCheckpoint := makeValsetCheckpoint(hyperionID, validators, powers, big.NewInt(0))

					err = outAbi.Copy(&state_lastValsetCheckpoint, out)
					Ω(err).Should(BeNil())
					Ω(state_lastValsetCheckpoint).Should(Equal(offchainCheckpoint))
				})

				_ = Describe("ValsetUpdatedEvent", func() {
					var (
						valsetUpdatedEvent = wrappers.HyperionValsetUpdatedEvent{}
					)

					BeforeEach(func() {
						_, err := ContractDeployer.Logs(
							context.Background(),
							hyperionLogsOpts,
							deployTxHash,
							"ValsetUpdatedEvent",
							unpackValsetUpdatedEventTo(&valsetUpdatedEvent),
						)
						orFail(err)
					})

					It("Should have valid Valset parameters", func() {
						Ω(valsetUpdatedEvent.NewValsetNonce).ShouldNot(BeNil())
						Ω(valsetUpdatedEvent.NewValsetNonce.String()).Should(Equal("0"))
						Ω(valsetUpdatedEvent.Validators).Should(BeEquivalentTo(validators))
						Ω(valsetUpdatedEvent.Powers).Should(BeEquivalentTo(powers))
					})
				})
			})

			_ = Describe("Valset Update", func() {
				var (
					updateValsetTxHash common.Hash
					updateValsetErr    error
					signValsetErr      error

					newValidators        []common.Address
					newPowers            []*big.Int
					valsetCheckpointHash common.Hash

					sigsV []uint8
					sigsR []common.Hash
					sigsS []common.Hash

					state_lastValsetNonce *big.Int
					nextValsetNonce       *big.Int
				)

				BeforeEach(func() {
					out, outAbi, err := ContractDeployer.Call(context.Background(), hyperionCallOpts,
						"state_lastValsetNonce", noArgs,
					)
					Ω(err).Should(BeNil())
					err = outAbi.Copy(&state_lastValsetNonce, out)
					Ω(err).Should(BeNil())
				})

				_ = When("ValsetUpdate being submitted", func() {
					BeforeEach(func() {
						// don't recompute new checkpoint
						if updateValsetTxHash != zeroHash {
							return
						}

						newValidators = make([]common.Address, len(validators))
						for i := 0; i < len(validators); i++ {
							// simply reverse the validator set
							newValidators[i] = validators[len(validators)-1-i]
						}

						newPowers = make([]*big.Int, len(powers))
						for i := 0; i < len(powers); i++ {
							// simply reverse the powers set
							newPowers[i] = powers[len(powers)-1-i]
						}

						nextValsetNonce = new(big.Int).Add(state_lastValsetNonce, big.NewInt(1))

						valsetCheckpointHash = makeValsetCheckpoint(
							hyperionID,
							newValidators,
							newPowers,
							nextValsetNonce,
						)

						sigsV, sigsR, sigsS, signValsetErr = signDigest(
							valsetCheckpointHash, getSigningKeys(CosmosAccounts[:3]...)...)
						orFail(signValsetErr)
					})

					JustBeforeEach(func() {
						// don't resend the batch
						if updateValsetTxHash != zeroHash {
							return
						}

						updateValsetTxHash, _, updateValsetErr = ContractDeployer.Tx(context.Background(), hyperionTxOpts,
							"updateValset", withArgsFn(
								// The new version of the validator set
								newValidators,   // address[] memory _newValidators,
								newPowers,       // uint256[] memory _newPowers,
								nextValsetNonce, // uint256 _newValsetNonce,
								// The current validators that approve the change
								validators,            // address[] memory _currentValidators,
								powers,                // uint256[] memory _currentPowers,
								state_lastValsetNonce, // uint256 _currentValsetNonce,
								// These are arrays of the parts of the current validator's signatures
								sigsV, // uint8[] memory _v,
								sigsR, // bytes32[] memory _r,
								sigsS, // bytes32[] memory _s
							))

					})

					_ = When("ValsetUpdate submission failed", func() {
						BeforeEach(func() {})
						AfterEach(func() {
							updateValsetTxHash = zeroHash
						})
					})

					_ = Context("ValsetUpdate submitted successfully", func() {
						BeforeEach(func() {
							orFail(updateValsetErr)
						})

						It("Updates Valset Nonce", func() {
							out, outAbi, err := ContractDeployer.Call(context.Background(), hyperionCallOpts,
								"state_lastValsetNonce", noArgs,
							)
							Ω(err).Should(BeNil())
							err = outAbi.Copy(&state_lastValsetNonce, out)
							Ω(err).Should(BeNil())

							Ω(state_lastValsetNonce).ShouldNot(BeNil())
							Ω(state_lastValsetNonce.String()).Should(Equal(nextValsetNonce.String()))
						})

						It("Updates Valset Checkpoint", func() {
							var state_lastValsetCheckpoint common.Hash

							out, outAbi, err := ContractDeployer.Call(context.Background(), hyperionCallOpts,
								"state_lastValsetCheckpoint", noArgs,
							)
							Ω(err).Should(BeNil())

							err = outAbi.Copy(&state_lastValsetCheckpoint, out)
							Ω(err).Should(BeNil())
							Ω(state_lastValsetCheckpoint).Should(Equal(valsetCheckpointHash))
						})

						_ = Describe("ValsetUpdatedEvent", func() {
							var (
								valsetUpdatedEvent = wrappers.HyperionValsetUpdatedEvent{}
							)

							BeforeEach(func() {
								_, err := ContractDeployer.Logs(
									context.Background(),
									hyperionLogsOpts,
									updateValsetTxHash,
									"ValsetUpdatedEvent",
									unpackValsetUpdatedEventTo(&valsetUpdatedEvent),
								)
								orFail(err)
							})

							It("Should have valid Valset parameters", func() {
								Ω(valsetUpdatedEvent.NewValsetNonce).ShouldNot(BeNil())
								Ω(valsetUpdatedEvent.NewValsetNonce.String()).Should(Equal(nextValsetNonce.String()))
								Ω(valsetUpdatedEvent.Validators).Should(BeEquivalentTo(newValidators))
								Ω(valsetUpdatedEvent.Powers).Should(BeEquivalentTo(newPowers))
							})
						})
					})
				})

				_ = When("ValsetUpdate submitted to rollback the Valset", func() {
					// NOTE: we could just override "validators" with "newValidator", as well as
					// "powers" with "newPowers", but instead we're going to rollback the Valset on the contract.

					var (
						rollbackValsetTxHash common.Hash
						rollbackValsetErr    error
					)

					BeforeEach(func() {
						if rollbackValsetTxHash != zeroHash {
							return
						}

						nextValsetNonce = new(big.Int).Add(state_lastValsetNonce, big.NewInt(1))

						valsetCheckpointHash = makeValsetCheckpoint(
							hyperionID,
							validators,
							powers,
							nextValsetNonce,
						)

						sigsV, sigsR, sigsS, signValsetErr = signDigest(
							valsetCheckpointHash, getSigningKeysForAddresses(newValidators, CosmosAccounts[:3]...)...)
						orFail(signValsetErr)

						// NOTE: this is a rollback, the current valset was the "new valset".
						rollbackValsetTxHash, _, rollbackValsetErr = ContractDeployer.Tx(context.Background(), hyperionTxOpts,
							"updateValset", withArgsFn(
								// The new version of the validator set
								validators,      // address[] memory _newValidators,
								powers,          // uint256[] memory _newPowers,
								nextValsetNonce, // uint256 _newValsetNonce,
								// The current validators that approve the change
								newValidators,         // address[] memory _currentValidators,
								newPowers,             // uint256[] memory _currentPowers,
								state_lastValsetNonce, // uint256 _currentValsetNonce,
								// These are arrays of the parts of the current validator's signatures
								sigsV, // uint8[] memory _v,
								sigsR, // bytes32[] memory _r,
								sigsS, // bytes32[] memory _s
							))
						orFail(rollbackValsetErr)
					})

					It("Updates Valset Nonce", func() {
						out, outAbi, err := ContractDeployer.Call(context.Background(), hyperionCallOpts,
							"state_lastValsetNonce", noArgs,
						)
						Ω(err).Should(BeNil())
						err = outAbi.Copy(&state_lastValsetNonce, out)
						Ω(err).Should(BeNil())

						Ω(state_lastValsetNonce).ShouldNot(BeNil())
						Ω(state_lastValsetNonce.String()).Should(Equal(nextValsetNonce.String()))
					})

					It("Updates Valset Checkpoint", func() {
						var state_lastValsetCheckpoint common.Hash

						out, outAbi, err := ContractDeployer.Call(context.Background(), hyperionCallOpts,
							"state_lastValsetCheckpoint", noArgs,
						)
						Ω(err).Should(BeNil())

						err = outAbi.Copy(&state_lastValsetCheckpoint, out)
						Ω(err).Should(BeNil())
						Ω(state_lastValsetCheckpoint).Should(Equal(valsetCheckpointHash))
					})

					_ = Describe("ValsetUpdatedEvent", func() {
						var (
							valsetUpdatedEvent = wrappers.HyperionValsetUpdatedEvent{}
						)

						BeforeEach(func() {
							_, err := ContractDeployer.Logs(
								context.Background(),
								hyperionLogsOpts,
								rollbackValsetTxHash,
								"ValsetUpdatedEvent",
								unpackValsetUpdatedEventTo(&valsetUpdatedEvent),
							)
							orFail(err)
						})

						It("Should have valid Valset parameters", func() {
							Ω(valsetUpdatedEvent.NewValsetNonce).ShouldNot(BeNil())
							Ω(valsetUpdatedEvent.NewValsetNonce.String()).Should(Equal(nextValsetNonce.String()))
							Ω(valsetUpdatedEvent.Validators).Should(BeEquivalentTo(validators))
							Ω(valsetUpdatedEvent.Powers).Should(BeEquivalentTo(powers))
						})
					})

				})
			})

			_ = Describe("ERC20 token deployment via Hyperion", func() {
				var (
					state_lastValsetNonce *big.Int
					state_lastEventNonce  *big.Int
					prevEventNonce        *big.Int

					erc20DeployTxHash  common.Hash
					erc20DeployErr     error
					erc20DeployedEvent = wrappers.HyperionERC20DeployedEvent{}
				)

				BeforeEach(func() {
					if state_lastEventNonce != nil {
						prevEventNonce = state_lastEventNonce
					}

					out, outAbi, err := ContractDeployer.Call(context.Background(), hyperionCallOpts,
						"state_lastValsetNonce", noArgs,
					)
					Ω(err).Should(BeNil())
					err = outAbi.Copy(&state_lastValsetNonce, out)
					Ω(err).Should(BeNil())
				})

				BeforeEach(func() {
					out, outAbi, err := ContractDeployer.Call(context.Background(), hyperionCallOpts,
						"state_lastEventNonce", noArgs,
					)
					Ω(err).Should(BeNil())
					err = outAbi.Copy(&state_lastEventNonce, out)
					Ω(err).Should(BeNil())
				})

				JustBeforeEach(func() {
					// don't redeploy
					if erc20DeployTxHash != zeroHash {
						return
					}

					erc20DeployTxHash, _, erc20DeployErr = ContractDeployer.Tx(context.Background(), hyperionTxOpts,
						"deployERC20", withArgsFn("helios", "HELIOS", "HELIOS", byte(18)),
					)
					orFail(erc20DeployErr)
				})

				It("Deploys new ERC20 contract instance", func() {
					Ω(erc20DeployErr).Should(BeNil())
					Ω(erc20DeployTxHash).ShouldNot(Equal(zeroHash))
				})

				_ = When("New ERC20 instance deployed successfully", func() {
					BeforeEach(func() {
						orFail(erc20DeployErr)

						// don't re-read the event
						if erc20DeployedEvent.TokenContract != zeroAddress {
							return
						}

						_, err := ContractDeployer.Logs(
							context.Background(),
							hyperionLogsOpts,
							erc20DeployTxHash,
							"ERC20DeployedEvent",
							unpackERC20DeployedEventTo(&erc20DeployedEvent),
						)
						orFail(err)
					})

					It("Nonce during deployment increased", func() {
						next := new(big.Int).Add(prevEventNonce, big.NewInt(1))
						Ω(state_lastEventNonce.String()).Should(Equal(next.String()))
					})

					_ = Describe("ERC20DeployedEvent", func() {
						It("Should have valid token params", func() {
							Ω(erc20DeployedEvent.CosmosDenom).Should(Equal("helios"))
							Ω(erc20DeployedEvent.Symbol).Should(Equal("HELIOS"))
							Ω(erc20DeployedEvent.Name).Should(Equal("HELIOS"))
							Ω(erc20DeployedEvent.Decimals).Should(BeEquivalentTo(18))
						})

						It("Should have TokenContract address", func() {
							Ω(erc20DeployedEvent.TokenContract).ShouldNot(Equal(zeroAddress))
						})

						It("Should have valid EventNonce", func() {
							Ω(erc20DeployedEvent.EventNonce).ShouldNot(BeNil())
							Ω(erc20DeployedEvent.EventNonce.String()).Should(Equal(state_lastEventNonce.String()))
						})
					})

					_ = Describe("ERC20 Token", func() {
						var (
							erc20CallOpts deployer.ContractCallOpts
						)

						BeforeEach(func() {
							erc20CallOpts = deployer.ContractCallOpts{
								From:          hyperionOwner.EthAddress,
								SolSource:     "../../../Ethereum-Bridge-Contract/contracts/CosmosToken.sol",
								ContractName:  "CosmosERC20",
								Contract:      erc20DeployedEvent.TokenContract,
								CoverageAgent: CoverageAgent,
								CoverageCall: deployer.ContractCoverageCallOpts{
									FromPk: hyperionOwner.EthPrivKey,
								},
							}
						})

						It("Should have MAX_UINT balance on Hyperion", func() {
							var hyperionBalance *big.Int

							out, outAbi, err := ContractDeployer.Call(context.Background(), erc20CallOpts,
								"balanceOf", withArgsFn(hyperionContract.Address))
							Ω(err).Should(BeNil())

							err = outAbi.Copy(&hyperionBalance, out)
							Ω(err).Should(BeNil())

							Ω(hyperionBalance).Should(BeEquivalentTo(maxUInt256()))
						})

						_ = When("Cosmos -> Ethereum batch being submitted", func() {
							var (
								submitBatchTxHash common.Hash
								submitBatchErr    error
								signBatchErr      error

								txAmounts      []*big.Int
								txDestinations []common.Address
								txFees         []*big.Int

								sigsV []uint8
								sigsR []common.Hash
								sigsS []common.Hash

								batchNonce   *big.Int
								batchTimeout *big.Int
							)

							BeforeEach(func() {
								// don't recompute hash
								if submitBatchTxHash != zeroHash {
									return
								}

								batchNonce = big.NewInt(1)
								batchTimeout = big.NewInt(10000)

								txAmounts = make([]*big.Int, len(EthAccounts))
								txDestinations = getEthAddresses(EthAccounts...)
								txFees = make([]*big.Int, len(EthAccounts))

								for i := range EthAccounts {
									txAmounts[i] = big.NewInt(1)
									txFees[i] = big.NewInt(1)
								}

								transactionBatchHash := prepareOutgoingTransferBatch(
									hyperionID,
									erc20DeployedEvent.TokenContract,
									txAmounts,
									txDestinations,
									txFees,
									batchNonce,
									batchTimeout,
								)

								sigsV, sigsR, sigsS, signBatchErr = signDigest(
									transactionBatchHash, getSigningKeys(CosmosAccounts[:3]...)...)
								orFail(signBatchErr)
							})

							JustBeforeEach(func() {
								// don't resend the batch
								if submitBatchTxHash != zeroHash {
									return
								}

								submitBatchTxHash, _, submitBatchErr = ContractDeployer.Tx(context.Background(), hyperionTxOpts,
									"submitBatch", withArgsFn(
										// The validators that approve the batch
										validators,            // 	address[] memory _currentValidators,
										powers,                // 	uint256[] memory _currentPowers,
										state_lastValsetNonce, // 	uint256 _currentValsetNonce,

										// These are arrays of the parts of the validators signatures
										sigsV, // 	uint8[] memory _v,
										sigsR, // 	bytes32[] memory _r,
										sigsS, // 	bytes32[] memory _s,

										// The batch of transactions
										txAmounts,                        // 	uint256[] memory _amounts,
										txDestinations,                   // 	address[] memory _destinations,
										txFees,                           // 	uint256[] memory _fees,
										batchNonce,                       // 	uint256 _batchNonce,
										erc20DeployedEvent.TokenContract, // 	address _tokenContract,

										// a block height beyond which this batch is not valid
										// used to provide a fee-free timeout
										batchTimeout, // 	uint256 _batchTimeout
									))

							})

							_ = When("TxBatch submission failed", func() {
								BeforeEach(func() {})
								AfterEach(func() {
									submitBatchTxHash = zeroHash
								})
							})

							_ = Context("TxBatch submitted successfully", func() {
								BeforeEach(func() {
									orFail(submitBatchErr)
								})

								It("Changes the balance of the Hyperion contract", func() {
									var hyperionBalance *big.Int

									out, outAbi, err := ContractDeployer.Call(context.Background(), erc20CallOpts,
										"balanceOf", withArgsFn(hyperionContract.Address))
									Ω(err).Should(BeNil())

									err = outAbi.Copy(&hyperionBalance, out)
									Ω(err).Should(BeNil())

									expenses := sumInts(nil, txAmounts...)
									expenses = sumInts(expenses, txFees...)
									remainder := new(big.Int).Sub(maxUInt256(), expenses)
									Ω(hyperionBalance.String()).Should(Equal(remainder.String()))
								})

								It("Increases the token balances of recipients", func() {
									recipients := getEthAddresses(EthAccounts...)
									for _, recipient := range recipients {
										var recipientBalance *big.Int

										out, outAbi, err := ContractDeployer.Call(context.Background(), erc20CallOpts,
											"balanceOf", withArgsFn(recipient))
										Ω(err).Should(BeNil())

										err = outAbi.Copy(&recipientBalance, out)
										Ω(err).Should(BeNil())

										if recipient == hyperionOwner.EthAddress {
											// the hyperionOwner address collected all the fees also
											Ω(recipientBalance.String()).Should(Equal(
												sumInts(big.NewInt(1), txFees...).String(),
											))
										} else {
											Ω(recipientBalance.String()).Should(Equal("1"))
										}
									}
								})

								It("Updates batch nonce for the token contract", func() {
									var lastBatchNonce *big.Int

									out, outAbi, err := ContractDeployer.Call(context.Background(), hyperionCallOpts,
										"lastBatchNonce", withArgsFn(erc20DeployedEvent.TokenContract))
									Ω(err).Should(BeNil())

									err = outAbi.Copy(&lastBatchNonce, out)
									Ω(err).Should(BeNil())

									Ω(lastBatchNonce).ShouldNot(BeNil())
									Ω(lastBatchNonce.String()).Should(Equal(batchNonce.String()))
								})

								_ = Describe("TransactionBatchExecutedEvent", func() {
									var (
										transactionBatchExecutedEvent = wrappers.HyperionTransactionBatchExecutedEvent{}
									)

									BeforeEach(func() {
										orFail(submitBatchErr)

										_, err := ContractDeployer.Logs(
											context.Background(),
											hyperionLogsOpts,
											submitBatchTxHash,
											"TransactionBatchExecutedEvent",
											unpackTransactionBatchExecutedEventTo(&transactionBatchExecutedEvent),
										)
										orFail(err)
									})

									It("Should have valid batch and nonce params", func() {
										Ω(transactionBatchExecutedEvent.EventNonce).ShouldNot(BeNil())
										Ω(transactionBatchExecutedEvent.EventNonce.String()).Should(Equal(state_lastEventNonce.String()))
										Ω(transactionBatchExecutedEvent.BatchNonce).ShouldNot(BeNil())
										Ω(transactionBatchExecutedEvent.BatchNonce.String()).Should(Equal(batchNonce.String()))
										Ω(transactionBatchExecutedEvent.Token).Should(Equal(erc20DeployedEvent.TokenContract))
									})
								})
							})
						})

						_ = When("Ethereum -> Cosmos batch being submitted", func() {

						})
					})
				})
			})
		})
	})
})

var outgoingBatchTxConfirmABI, _ = abi.JSON(strings.NewReader(hyperion.OutgoingBatchTxConfirmABIJSON))

func prepareOutgoingTransferBatch(
	hyperionID common.Hash,
	tokenContract common.Address,
	txAmounts []*big.Int,
	txDestinations []common.Address,
	txFees []*big.Int,
	batchNonce *big.Int,
	batchTimeout *big.Int,
) common.Hash {
	abiEncodedBatch, err := outgoingBatchTxConfirmABI.Pack("transactionBatch",
		hyperionID,
		formatBytes32String("transactionBatch"),
		txAmounts,
		txDestinations,
		txFees,
		batchNonce,
		tokenContract,
		batchTimeout,
	)
	orFail(err)

	return crypto.Keccak256Hash(abiEncodedBatch[4:])
}

func unpackERC20DeployedEventTo(result *wrappers.HyperionERC20DeployedEvent) deployer.ContractLogUnpackFunc {
	return func(unpacker deployer.LogUnpacker, event abi.Event, log ctypes.Log) (interface{}, error) {
		err := unpacker.UnpackLog(result, event.Name, log)
		return &result, err
	}
}

func unpackValsetUpdatedEventTo(result *wrappers.HyperionValsetUpdatedEvent) deployer.ContractLogUnpackFunc {
	return func(unpacker deployer.LogUnpacker, event abi.Event, log ctypes.Log) (interface{}, error) {
		err := unpacker.UnpackLog(result, event.Name, log)
		return &result, err
	}
}

func unpackTransactionBatchExecutedEventTo(result *wrappers.HyperionTransactionBatchExecutedEvent) deployer.ContractLogUnpackFunc {
	return func(unpacker deployer.LogUnpacker, event abi.Event, log ctypes.Log) (interface{}, error) {
		err := unpacker.UnpackLog(result, event.Name, log)
		return &result, err
	}
}
