const fs = require('fs');
const path = require('path');
const { gql } = require('apollo-server');

const schema = fs.readFileSync(
  path.join(__dirname, '../../../api/infrastructure/graph/schema.graphqls'),
  'utf8'
);
const typeDefs = gql(schema);

module.exports = { typeDefs };
