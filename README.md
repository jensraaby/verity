Verity
========

I have a problem: I have thousands of binary files (photos/music) which I do not want to one day become corrupt. I have several backup strategies, but what happens if the original file is silently corrupted and then replicated to all the backups?

This project is not even alpha status yet, as I am playing around a lot with how to architect the solution.
It could be done in many trivial ways (even shell scripting) but I wanted a small project to play with Google's Go language. Here are a few of the ideas I've had for the implementation:

Ideas
=====
- Create a pipeline using channels to pass data:
  - 1. scan file system for directories, pushing these onto a channel
  - 2. spin up a parallel goroutine for each directory
  - 3. create a pool of goroutines to perform SHA-1 hashing
  - 4. send files in a directory to the workers until there are no files remaining
  - 5. serialise each directory's hashes to a hidden file as a small JSON file.
- Extend to handle comparing two directory trees
  - Extend to run on multiple hosts (passing computed hashes back over the network for comparison)
- Possibly serialise the file in the root directory rather than in each separate directory



Please don't fork this repo yet!
