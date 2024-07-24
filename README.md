# The Problem

The team have written their first Go service; a simple user processor that creates, gets, and deletes a user. The problem is, they can't write unit tests for the critical `ProcessUserRequest` function because it requires an running database.

# The Challenge

The main point of this challenge is to refactor the code in a idiomatic way, such that the `ProcessUserRequest` function can be unit tested.

There are a few other gotchas in the code, bonus points if you can spot them.
