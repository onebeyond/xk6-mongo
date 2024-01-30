import xk6_mongo from 'k6/x/mongo';


const client = xk6_mongo.newClient('mongodb://localhost:27017', 'testcollection');
export default ()=> {

    let doc = {
        correlationId: `test--mongodb`,
        title: 'Perf test experiment',
        url: 'example.com',
        locale: 'en',
        time: `${new Date(Date.now()).toISOString()}`
      };

      client.insertOne("testdb", doc);
}
