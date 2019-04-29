#!/usr/bin/env python3

import os, time, subprocess

def runCmd(cmd):
    p = subprocess.Popen(cmd, shell = True, stdout = subprocess.PIPE, stderr = subprocess.PIPE)
    stdout = p.communicate()[0].decode('utf-8').strip()
    return stdout
def modPath():
    return 'APIServer/config'

# Get last tag.
def lastTag():
    return runCmd('git describe --abbrev=0 --tags')

# Get current branch name.
def branch():
    return runCmd('git rev-parse --abbrev-ref HEAD')

# Get last git commit id.
def lastCommitId():
    return runCmd('git log --pretty=format:"%h" -1')

# Assemble build command.
def buildCmd():
    buildFlag = []

    modName = modPath()
    currentTime = time.strftime("%Y-%m-%d-%H:%M")

    version = lastTag()
    if version != "":
        buildFlag.append("-X {}._version_='{}'".format(modName, version))

    branchName = branch()
    if branchName != "":
        buildFlag.append("-X {}._branch_='{}'".format(modName, branchName))

    commitId = lastCommitId()
    if commitId != "":
        buildFlag.append("-X {}._commitId_='{}'".format(modName, commitId))

        # current time
        buildFlag.append("-X {}._buildTime_='{}'".format(modName, currentTime))
        
    return 'go build -ldflags "{}" -o ./bin/main-{}.exe ./main.go'.format(" ".join(buildFlag), version)


if __name__ == "__main__":
    print(buildCmd())
    
    if subprocess.call(buildCmd(), shell = True) == 0:
        print("build finished.")