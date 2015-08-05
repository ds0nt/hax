COND := var i = 100; while (i-- > 0)
EMAIL := echo "$x $x"

all: pre email
.PHONY: all

pre: node_modules
	@ npm install --save faker
.PHONY: pre

email:
	@ iojs -e "var fn = require('faker').internet.email; ${COND} { console.log(fn()); }" | echo $(read)
.PHONY: email

# var randomName = faker.name.findName();
# var randomEmail = faker.internet.email();
# var randomCard = faker.helpers.createCard();

curl 'http://appscend.us2.list-manage.com/subscribe/post?u=de3298de9232b91872f2c5529&amp;id=2cf33c9443' -H 'Pragma: no-cache' -H 'Origin: http://docs.telepat.io' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-US,en;q=0.8,ja;q=0.6,zh-CN;q=0.4' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.81 Safari/537.36' -H 'Content-Type: application/x-www-form-urlencoded' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8' -H 'Cache-Control: no-cache' -H 'Referer: http://docs.telepat.io/intro.html' -H 'Cookie: __utma=239944220.1023369494.1435133560.1435133560.1435133560.1; __utmc=239944220; __utmz=239944220.1435133560.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none)' -H 'Connection: keep-alive' --data 'b_de3298de9232b91872f2c5529_2cf33c9443=&EMAIL=open.source.or.gtfo%40gmail.com&subscribe=' --compressed
