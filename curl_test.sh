#!/bin/bash

T_OK='47e9b9c8ffb7cc4310f72189ef9d5cc9cf56d1c286083f5e41118c1f1b91c273'
T_VOIP='da02c395db49b9918706b34a7dfba7cb670120b6d254be41572c1f960498848a'
T_FAIL='05599d8ea99c2c924c639c13873393d79abbf29453db952933ed11adeba0b8d8'


/usr/bin/curl -1kv --http2 --http2-prior-knowledge \
-H "method : POST'" \
-H "path : /3/device/$T_OK" \
-H "apns-push-type : alert" \
-H "apns-topic : com.mitake.mitakeeim" \
-d "{\"aps\":{\"alert\":\"Hello\"}}" \
-X POST https://api.push.apple.com
