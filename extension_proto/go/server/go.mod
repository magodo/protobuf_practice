module server

require (
	git.ucloudadmin.com/udb/v2/common v0.0.0-20190117055312-70cc0e06f9ad
	git.ucloudadmin.com/udb/v2/udb_access2 v0.0.0-20190125025444-9dce417bec5a
	github.com/pkg/errors v0.8.1
	gitlab.ucloudadmin.com/udb/uframework v0.0.0-20181204102014-865522b572fd
	proto_foo v0.0.0
)

replace proto_foo => /home/magodo/github/protobuf_practice/extension_proto/go/proto-mod