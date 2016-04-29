# Gitaxis

**A proof of concept of a concurrent and distributed botnet written in Go.**
**Each Gitaxis node consists of 3, possibly more, microservices (See: microservices)**

**__This project just started development, Keeper is the only thing in a functioning state at the moment__**

A user of an instance of the Gitaxis botnet would be able to log onto any of the nodes. There's no centralized 'master' of 'slaves'. In order to have any power over the botnet, the only necessity is for the node to query other nodes on the network before validating the instruction as trusted.


### Current Issue:
   * Currently trying to design a way for node password protection (to decrypt stored data) to work.
   * We don't want users of the 'host' computer to be able access that information (i.e. ip's of other nodes)
   * It need's to be done in such a way that someone can't just decompile the code to find the password, and it needs to be easily changed
   * One possible solution is the idea of accounts that are local to individual nodes, but having account information be stored on other nodes, but I'm not sure if that'd work

## Microservices:

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
