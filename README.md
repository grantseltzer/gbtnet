# Gitaxis 

###### - A proof of concept of a concurrent and decentralized botnet written in Go.
###### - Each Gitaxis node consists of 3 microservices: Keeper, Membrane, and Whisper.

A user of an instance of the Gitaxis botnet would be able to log onto any of the nodes. There's not centralized 'master' of slaves. In order to have any power over the botnet, the only necessity is for the node to query other nodes on the network before validating the instruction as trusted.

#### Membrane -
    * Public facing REST api
    * Receives instructions/queries from other nodes on the botnet
    * Filters, examines, and controls incoming traffic.
    * Delegates instructions to Keeper

#### Keeper -
    * Local REST api
    * Receives instructions from Membrane
    * Interacts with stored information (i.e. list of ip's on blockchain)
    * Queue's instructions for Whisper to pick up and execute

#### Whisper - 
    * Local REST api
    * Picks up instructions to execute from Keeper
    * Sends out information to other nodes
    * Executes payloads
    
