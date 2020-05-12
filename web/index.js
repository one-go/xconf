const {ListConfigsRequest} = require('./xconf_pb.js');
const {XconfClient} = require('./xconf_grpc_web_pb.js');

var client = new XconfClient('http://127.0.0.1:8080', null, null);

var request = new ListConfigsRequest();
request.setParent('ifish/rio');

client.listConfigs(request, {}, (err, response) => {
  console.log(err, response.getConfigsList());
});
