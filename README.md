GoRobot
===

## an IRC bot written in Go

## Features:

  * Multiple servers
  * Multiple channels
  * Conversations
  * Administration via a specific set of IRC channels
  * Go modules, which can connect to the bot through netchans (one can dynamically add new modules through network)
  * Simple API to write GO modules
  * JSON Configuration
  * No excess flood, even upon high activity
  * Module to handle shell scripts, through a tiny API
  * Module to Broadcast conversations on a specific IRC channel

## Todo:

  * Statistics (activity on a channel, number of people, ...)
  * A prompt to perform administrative tasks?
  * Automatic reconnection on timeout, kicks, ...
  * A GO module which acts as an IRC client? :-)
  * Add new features to the GO API (configuration of modules, ...)
  * Ability to autoload GO modules