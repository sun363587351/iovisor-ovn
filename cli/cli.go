// Copyright 2016 Politecnico di Torino
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Inline command line interface for debug purposes
// in future this cli will be a separate go program that connects to the main iovisor-ovn daemon
// in future this cli will use a cli go library (e.g. github.com/urfave/cli )

package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/netgroup-polito/iovisor-ovn/hoverctl"
	"github.com/netgroup-polito/iovisor-ovn/iomodules/l2switch"
	"github.com/netgroup-polito/iovisor-ovn/mainlogic"
	"github.com/netgroup-polito/iovisor-ovn/ovnmonitor"
)

func Cli(dataplaneref *hoverctl.Dataplane) {
	db := &mainlogic.Mon.DB
	dataplane := dataplaneref
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("cli@iov-ovn$ ")
		line, _ := reader.ReadString('\n')

		line = TrimSuffix(line, "\n")
		args := strings.Split(line, " ")

		if len(args) >= 1 {
			switch args[0] {
			case "mainlogic", "ml":
				if len(args) >= 2 {
					switch args[1] {
					case "-v":
						fmt.Printf("\nMainLogic (verbose)\n\n")
						mainlogic.PrintMainLogic(true)
					case "switch":
						if len(args) >= 3 {
							fmt.Printf("\nMainLogic Switch %s\n\n", args[2])
							mainlogic.PrintL2Switch(args[2])
						} else {
							fmt.Printf("\nMainLogic Switches \n\n")
							mainlogic.PrintL2Switches(true)
						}
					case "router":
						if len(args) >= 3 {
							fmt.Printf("\nMainLogic Router %s\n\n", args[2])
							mainlogic.PrintRouter(args[2])
						} else {
							fmt.Printf("\nMainLogic Routers \n\n")
							mainlogic.PrintRouters(true)
						}
					}
				} else {
					fmt.Printf("\nMainLogic\n\n")
					mainlogic.PrintMainLogic(false)
				}

			case "ovnmonitor", "ovn", "o":
				if len(args) >= 2 {
					switch args[1] {
					case "-v":
						fmt.Printf("\nOvnMonitor (verbose)\n\n")
						ovnmonitor.PrintOvnMonitor(true, db)
					case "switch", "s":
						if len(args) >= 3 {
							fmt.Printf("\nOvnMonitor Logical Switch %s\n\n", args[2])
							ovnmonitor.PrintLogicalSwitchByName(args[2], db)
						} else {
							fmt.Printf("\nOvnMonitor Logical Switches \n\n")
							ovnmonitor.PrintLogicalSwitches(true, db)
						}
					case "router", "r":
						if len(args) >= 3 {
							fmt.Printf("\nOvnMonitor Logical Router %s\n\n", args[2])
							ovnmonitor.PrintLogicalRouterByName(args[2], db)
						} else {
							fmt.Printf("\nOvnMonitor Logical Routers \n\n")
							ovnmonitor.PrintLogicalRouters(true, db)
						}
					case "interface", "i":
						fmt.Printf("\nOvnMonitor Ovs Interfaces\n\n")
						ovnmonitor.PrintOvsInterfaces(db)
					}
				} else {
					fmt.Printf("\nOvn Monitor\n\n")
					ovnmonitor.PrintOvnMonitor(false, db)
				}

			// case "test":
			// 	fmt.Printf("\ntest\n\n")
			// 	//tests.TestModule(dataplane)
			// case "config":
			// 	config.PrintConfigCli()
			// 	if len(args) == 3 {
			// 		switch args[1] {
			// 		case "PrintOvnNbChanges":
			// 			switch args[2] {
			// 			case "true":
			// 				config.PrintOvnNbChanges = true
			// 				fmt.Printf("OK\n")
			// 			case "false":
			// 				config.PrintOvnNbChanges = false
			// 				fmt.Printf("OK\n")
			// 			}
			// 		case "PrintOvnSbChanges":
			// 			switch args[2] {
			// 			case "true":
			// 				config.PrintOvnSbChanges = true
			// 				fmt.Printf("OK\n")
			// 			case "false":
			// 				config.PrintOvnSbChanges = false
			// 				fmt.Printf("OK\n")
			// 			}
			// 		case "PrintOvsChanges":
			// 			switch args[2] {
			// 			case "true":
			// 				config.PrintOvsChanges = true
			// 				fmt.Printf("OK\n")
			// 			case "false":
			// 				config.PrintOvsChanges = false
			// 				fmt.Printf("OK\n")
			// 			}
			// 		}
			// 	}
			// case "ovncontroller", "o":
			// 	if len(args) >= 1 {
			// 		if len(args) == 1 {
			// 			fmt.Printf("\n*********************OVN-NorthBound-Database***********************\n\n")
			// 			ovnmonitor.PrintCache(hh.Nb)
			// 			fmt.Printf("\n*********************OVN-SouthBound-Database***********************\n\n")
			// 			ovnmonitor.PrintCache(hh.Sb)
			// 			fmt.Printf("\n************************OVS-Local-Database*************************\n\n")
			// 			ovnmonitor.PrintCache(hh.Ovs)
			// 		}
			// 		if len(args) == 2 {
			// 			//switch
			// 			switch args[1] {
			// 			case "nb":
			// 				fmt.Printf("\n*********************OVN-NorthBound-Database***********************\n\n")
			// 				ovnmonitor.PrintCache(hh.Nb)
			// 			case "sb":
			// 				fmt.Printf("\n*********************OVN-SouthBound-Database***********************\n\n")
			// 				ovnmonitor.PrintCache(hh.Sb)
			// 			case "ovs":
			// 				fmt.Printf("\n************************OVS-Local-Database*************************\n\n")
			// 				ovnmonitor.PrintCache(hh.Ovs)
			// 			default:
			// 				PrintOvnControllerUsage()
			// 			}
			// 		}
			// 		if len(args) >= 3 {
			// 			PrintOvnControllerUsage()
			// 		}
			// 	} else {
			// 		PrintOvnControllerUsage()
			// 	}
			// case "ovs":
			// 	if len(args) >= 1 {
			// 		if len(args) == 1 {
			// 			ovnmonitor.PrintOvs(hh)
			// 		} else {
			// 			PrintOvsUsage()
			// 		}
			// 	} else {
			// 		PrintOvsUsage()
			// 	}
			// case "nb":
			// 	if len(args) >= 1 {
			// 		if len(args) == 1 {
			// 			ovnmonitor.PrintNb(hh)
			// 		} else {
			// 			if len(args) == 3 {
			// 				switch args[2] {
			// 				case "ls":
			// 					ovnmonitor.PrintNbLogicalSwitch(hh, args[2])
			// 				case "lsp":
			// 					ovnmonitor.PrintNbLogicalSwitchPort(hh, args[2])
			// 				default:
			// 					PrintNbUsage()
			// 				}
			// 			} else {
			// 				PrintNbUsage()
			// 			}
			// 		}
			// 	} else {
			// 		PrintNbUsage()
			// 	}
			// 	//fmt.Printf("\nNorthBound DB\n\n")
			// 	//ovnmonitor.PrintNb(hh)
			case "interfaces", "i":
				fmt.Printf("\nInterfaces\n\n")
				_, external_interfaces := hoverctl.ExternalInterfacesListGET(dataplane)
				hoverctl.ExternalInterfacesListPrint(external_interfaces)
			case "modules", "m":
				if len(args) >= 2 {
					switch args[1] {
					case "get":
						switch len(args) {
						case 2:
							fmt.Printf("\nModules GET\n\n")
							_, modules := hoverctl.ModuleListGET(dataplane)
							hoverctl.ModuleListPrint(modules)
						case 3:
							fmt.Printf("\nModules GET\n\n")
							_, module := hoverctl.ModuleGET(dataplane, args[2])
							hoverctl.ModulePrint(module)
						default:
							PrintModulesUsage()
						}
					case "post":
						switch len(args) {
						case 3:
							fmt.Printf("\nModules POST\n\n")
							if args[2] == "switch" {
								_, module := hoverctl.ModulePOST(dataplane, "bpf", "Switch", l2switch.SwitchSecurityPolicy)
								hoverctl.ModulePrint(module)
							} else {
								//TODO Print modules list
							}
						default:
							PrintModulesUsage()
						}
					case "delete":
						switch len(args) {
						case 3:
							fmt.Printf("\nModules DELETE\n\n")
							hoverctl.ModuleDELETE(dataplane, args[2])
						default:
							PrintModulesUsage()
						}
					default:
						PrintModulesUsage()
					}
				} else {
					PrintModulesUsage()
				}
			case "links", "l":
				if len(args) >= 2 {
					switch args[1] {
					case "get":
						switch len(args) {
						case 2:
							fmt.Printf("\nLinks GET\n\n")
							_, links := hoverctl.LinkListGet(dataplane)
							hoverctl.LinkListPrint(links)
						case 3:
							fmt.Printf("\nLinks GET\n\n")
							_, link := hoverctl.LinkGET(dataplane, args[2])
							hoverctl.LinkPrint(link)
						default:
							PrintLinksUsage()
						}
					case "post":
						switch len(args) {
						case 4:
							fmt.Printf("\nLinks POST\n\n")
							_, link := hoverctl.LinkPOST(dataplane, args[2], args[3])
							hoverctl.LinkPrint(link)
						default:
							PrintLinksUsage()
						}
					case "delete":
						switch len(args) {
						case 3:
							fmt.Printf("\nLinks DELETE\n\n")
							hoverctl.LinkDELETE(dataplane, args[2])
						default:
							PrintLinksUsage()
						}
					default:
						PrintLinksUsage()
					}
				} else {
					PrintLinksUsage()
				}
			case "table", "t":
				if len(args) >= 2 {
					switch args[1] {
					case "get":
						switch len(args) {
						case 2:
							fmt.Printf("\nTable GET\n\n")
							_, modules := hoverctl.ModuleListGET(dataplane)
							for moduleName, _ := range modules {
								fmt.Printf("**MODULE** -> %s\n", moduleName)
								_, tables := hoverctl.TableListGET(dataplane, moduleName)
								for _, tablename := range tables {
									fmt.Printf("Table *%s*\n", tablename)
									_, table := hoverctl.TableGET(dataplane, moduleName, tablename)
									hoverctl.TablePrint(table)
								}
							}
						case 3:
							fmt.Printf("\nTable GET\n\n")
							_, tables := hoverctl.TableListGET(dataplane, args[2])
							for _, tablename := range tables {
								fmt.Printf("Table *%s*\n", tablename)
								_, table := hoverctl.TableGET(dataplane, args[2], tablename)
								hoverctl.TablePrint(table)
							}
						case 4:
							fmt.Printf("\nTable GET\n\n")
							_, table := hoverctl.TableGET(dataplane, args[2], args[3])
							hoverctl.TablePrint(table)
						case 5:
							fmt.Printf("\nTable GET\n\n")
							_, tableEntry := hoverctl.TableEntryGET(dataplane, args[2], args[3], args[4])
							hoverctl.TableEntryPrint(tableEntry)
						default:
							PrintTableUsage()
						}
					case "put":
						if len(args) == 6 {
							fmt.Printf("\nTable PUT\n\n")
							_, tableEntry := hoverctl.TableEntryPUT(dataplane, args[2], args[3], args[4], args[5])
							hoverctl.TableEntryPrint(tableEntry)
						} else {
							PrintTableUsage()
						}
					case "delete":
						if len(args) == 5 {
							fmt.Printf("\nTable DELETE\n\n")
							hoverctl.TableEntryDELETE(dataplane, args[2], args[3], args[4])
						} else {
							PrintTableUsage()
						}
					default:
						PrintTableUsage()
					}
				} else {
					PrintTableUsage()
				}
			case "help", "h":
				PrintHelp()

			case "":
			default:
				fmt.Printf("\nInvalid Command\n\n")
				PrintHelp()
			}
		}
		fmt.Printf("\n")
	}
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

//
// func PrintConfigUsage() {
// 	fmt.Printf("\nConfig Usage\n\n")
// 	fmt.Printf("	config              print Config\n")
// 	fmt.Printf("	config <parameter> <value>\n")
// }
//
// func PrintOvnControllerUsage() {
// 	fmt.Printf("\nOvn Controller Usage\n\n")
// 	fmt.Printf("	ovncontroller       print Databases\n")
// 	fmt.Printf("	ovncontroller nb		print Nb\n")
// 	fmt.Printf("	ovncontroller sb		print Sb\n")
// 	fmt.Printf("	ovncontroller ovs		print Ovs\n")
// }
//
// func PrintOvsUsage() {
// 	fmt.Printf("\nOVS Usage\n\n")
// 	fmt.Printf("	ovs     print the whole Ovs Local Database\n")
// }
//
// func PrintNbUsage() {
// 	fmt.Printf("\nNB Usage\n\n")
// 	fmt.Printf("	nb                  print the whole NorthBound\n")
// 	fmt.Printf("	nb ls   <ls-name>   print the Logical Switch table\n")
// 	fmt.Printf("	nb lsp  <lsp-name>  print the Logical Switch Port table\n")
// }

func PrintTableUsage() {
	fmt.Printf("\nTable Usage\n\n")
	fmt.Printf("	table get\n")
	fmt.Printf("	table get <module-id>\n")
	fmt.Printf("	table get <module-id> <table-id>\n")
	fmt.Printf("	table get <module-id> <table-id> <entry-key>\n")
	fmt.Printf("	table put <module-id> <table-id> <entry-key> <entry-value>\n")
	fmt.Printf("	table delete <module-id> <table-id> <entry-key> <entry-value>\n")
}

func PrintLinksUsage() {
	fmt.Printf("\nLinks Usage\n\n")
	fmt.Printf("	links get\n")
	fmt.Printf("	links get <link-id>\n")
	fmt.Printf("	links post <from> <to>\n")
	fmt.Printf("	links delete <link-id>\n")
}

func PrintModulesUsage() {
	fmt.Printf("\nModules Usage\n\n")
	fmt.Printf("	modules get\n")
	fmt.Printf("	modules get <module-id>\n")
	fmt.Printf("	modules post <module-name>\n")
	fmt.Printf("	modules delete <module-id>\n")
}

func PrintMainLogicUsage() {
	fmt.Printf("\nMainLogic Usage\n\n")
	fmt.Printf("	mainlogic (-v)\n")
	fmt.Printf("	mainlogic switch (<switch-name>)\n")
	fmt.Printf("	mainlogic router (<router-name>)\n")
}

func PrintOvnMonitorUsage() {
	fmt.Printf("\nOvnMonitor Usage\n\n")
	fmt.Printf("	ovnmonitor (-v)\n")
	fmt.Printf("	ovnmonitor switch (<switch-name>)\n")
	fmt.Printf("	ovnmonitor router (<router-name>)\n")
	fmt.Printf("	ovnmonitor interface\n")
}

func PrintHelp() {
	fmt.Printf("\n")
	fmt.Printf("IOVisor-OVN Command Line Interface HELP\n\n")
	fmt.Printf("	interfaces, i    prints /external_interfaces/\n")
	fmt.Printf("	modules, m       prints /modules/\n")
	fmt.Printf("	links, l         prints /links/\n")
	fmt.Printf("	table, t         prints tables\n\n")
	fmt.Printf("	mainlogic, ml    prints mainlogic\n")
	fmt.Printf("	ovnmonitor, ovn  prints ovnmonitor\n\n")

	// fmt.Printf("	nb               prints NorthBound database local structs\n")
	// fmt.Printf("	ovs              prints Ovs local database local structs\n\n")
	// fmt.Printf("	config,c         config print and modify\n\n")
	// fmt.Printf("	ovncontroller,o  prints OVN Databases\n\n")
	fmt.Printf("	help, h          print help\n")
	fmt.Printf("\n")
	PrintModulesUsage()
	PrintLinksUsage()
	PrintTableUsage()
	PrintMainLogicUsage()
	PrintOvnMonitorUsage()
	// PrintNbUsage()
	// PrintOvsUsage()
	// PrintConfigUsage()
	// PrintOvnControllerUsage()
}
