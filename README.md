# Gitaxis
[![forthebadge](http://forthebadge.com/images/badges/contains-cat-gifs.svg)](http://forthebadge.com)

**A proof of concept of a concurrent and distributed botnet written in Go.**

**__This project just started development__**

Please take up an issue on the repo if you have any questions!

A user of an instance of the Gitaxis botnet would be able to log onto any of the nodes. There's no centralized 'master' of 'slaves'. The advantage being no single point of failure, and making the orchestrator that much harder to track. In order to have any power over the botnet, the only necessity is for the node to query other nodes on the network before validating the instruction as trusted.

## Key Concepts:
- A node on any computer can be interacted with by sending HTTP requests from any other computer (with or without Gitaxis on it)
- There is secure blockchain-inspired validation to make sure the botnet owner is the one logging into any of the nodes on the botnet network
- Nodes can resiliently rejoin the botnet even if the host computer goes down for a non-lengthy amount of time
- No single point of detection, and no single point of failure
- 100% Golang, 100% badass

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
