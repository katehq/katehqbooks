#!/bin/bash
ps -ef | grep ./katehqbooks | grep -v grep | awk '{print $2}' | xargs kill -9