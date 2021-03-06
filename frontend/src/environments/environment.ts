// The file contents for the current environment will overwrite these during build.
// The build system defaults to the dev environment which uses `environment.ts`, but if you do
// `ng build --env=prod` then `environment.prod.ts` will be used instead.
// The list of which env maps to which file can be found in `angular-cli.json`.

export const environment = {
  production: false,
  laravel: '/api',
  go:'http://0.0.0.0:3002/api',
  go_prods:'http://0.0.0.0:3000/api/products',
  go_buy:'http://0.0.0.0:3001/api/buy_products',

};
