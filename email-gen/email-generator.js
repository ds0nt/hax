var faker = require('faker')

function* email() {
  var i = 0;
  while (true) {
    i++;
    yield i + ": " + faker.internet.email();
    if (i % 1000) {
      
    }
  }
}
var gen = email();
for (var next of gen) {
  console.log(next);

}