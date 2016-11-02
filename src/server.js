import koa from 'koa';
import log from './log';

const PORT = process.env.PORT || 8888;

const app = koa();

app.use(function *(next) {
  log.info('%s - %s', this.method, this.url);
  yield next;
});

log.info(`Listening on port ${PORT}`)
app.listen(PORT);