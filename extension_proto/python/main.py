#!/usr/bin/env python3
# -*- coding: utf-8 -*-

#########################################################################
# Author: Zhaoting Weng
# Created Time: Wed 11 Jul 2018 03:02:13 PM CST
# Description:
#########################################################################

import google.protobuf.text_format as tf
import foo_pb2, ext_foo_pb2

msg = foo_pb2.Message()

header = foo_pb2.Header()
header.type = 123
msg.header.CopyFrom(header)

body = foo_pb2.Body()
body.Extensions[ext_foo_pb2.name] = 'name'
body.Extensions[ext_foo_pb2.id] = 123
body.Extensions[ext_foo_pb2.wod] = ext_foo_pb2.Fri
msg.body.CopyFrom(body)

bin_msg = msg.SerializeToString()

msg2 = foo_pb2.Message()
msg2.ParseFromString(bin_msg)
print(ext_foo_pb2.WOD.Name(msg2.body.Extensions[ext_foo_pb2.wod]))
print(tf.MessageToString(msg2))
