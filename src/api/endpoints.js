import _ from 'underscore';
import koaRouter from 'koa-router';
import read from 'read-file';
import marked from 'marked';
import path from 'path';
import log from '../utils/log';

const router = koaRouter();
const README = 'README.md';

export default router;

router.get('/', function*(next) {
  this.body = marked(read.sync(README, 'utf8'));
});