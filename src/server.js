import koa from 'koa';
import log from './utils/log';
import endpoints from './api/endpoints';

const PORT = process.env.PORT || 8888;
const app = koa();



app.use(function *(next) {
  log.info('%s - %s', this.method, this.url);
  yield next;
});

app.use(endpoints.routes());
app.use(endpoints.allowedMethods());

log.info(`Listening on port ${PORT}`);
app.listen(PORT);