# lagerauth

mono repo like all the cool kids.

## projects:
**lagerauth** is the backend application, this also hosts the "compiled" site.  
**site** is the vue.js SPA frontend manager application.  
**clients** contains clients for lagerauth.  

check each project for their respective readme(s)

## docker deploy:
have a file named `prod.env.js` for site build and `conf.json` for lagerauth.
then you can run `docker build -t lagerauth .` no problem.

remember to bind the port you added to `bind` on `conf.json` for example:

`docker run --name lagerauth --rm -p 80:8081 lagerauth`


### notes on config and docker deploy:
if you change a file from `prod.env.js` or `conf.json` you will need to rebuild the image and re-start the container, since this files arent supposed to change that much i just copied them right on.

theres also the posibility of you running a read-only volume on `conf.json` without issues, but changing `prod.env.json` would require to re-build the site and therefore re-build the image.