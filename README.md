[WSU Hackathon Website](http://hackathon.eecs.wsu.edu)
======================

Build Instructions
------------------
#### Short Story
If you've installed and configured everything, generally all you'll need is these two commands:
```sh
npm install  ## Grab any missing dependencies
./serve.sh   ## Start the dev server at http://localhost:3030
```
#### Long Story
1. Install Node.js and npm. This will be dependant upon your system. Note that for
   those new to node on Debian systems, the nodejs-legacy package
   will probably be required. It's also probably worthwhile to update npm to
   the most latest version with:
   ```
   sudo npm install npm -g
   ```
   and then reload your shell.

2. Install gulp.
   ```
   sudo npm install gulp -g
   ```

3. Grab the dependencies. This may occasionally have to be rerun as we update
   the site and add additional dependencies.
   ```
   npm install
   ```

4. Configure the API server. Start with the file at
   ```configs/default-api-server-config.json``` and copy it to
   ```configs/api-server-config.json```. Fill in the Eventbrite eventId and
   oathToken params with suitable values. Whatever you do, do NOT add the
   Eventbrite credentials to the ```default-api-server-config.json``` file and
   proceed to commit it into the repo! After setting up the config, you'll need
   to grab some images to put in your hosted-images directory.
   This can be done simply enough by running
   ```
   ./seed-hosted-images.sh
   ```

5. Install Jekyll. If you already have ruby, and ruby gems installed,
   this can be as simple as running
   ```
   gem install jekyll
   ```
   If not, Jekyll may be in your distro's repositories
   (a big maybe as it will probably be an older version) or you can
   choose to take the recommended route and
   install ruby from scratch. Most people use [RVM](rvm.io) to do that.
   Whichever way you decide, the oldest version of
   Jekyll known to work is 2.5.3.

6. At this point, building and serving is as easy as calling
   ```
   ./serve.sh
   ```
   This script will build and copy the assets into a
   directory called ```build```. From there they will be served at the address
   [http://localhost:3030](http://localhost:3030). The build will automatically 
   be rerun as you make
   file changes and all you'll have to do is refresh your web page =)
   To only do a build with no file
   watching or serving, run ```jekyll build```.

***Bonus:***  
If you've downloaded the source for the
[api server](https://github.com/WSU-ACM/hackathon-api-server) you can have npm
directly link into it so you don't continually have to be pushing code around.
To do this, run the following in the directory of the hackathon-api-server source
```
sudo npm link
```
Then, in this repo run
```
npm link hackathon-api-server
```


Deploy Instructions
-------------------
So you've made your changes and think you're ready to deploy? If you have a user
account on the server, this is as easy as running ```./deploy.sh```. Just enter
your user name and it will do the rest. To make your life easier, it's
recommended to have public key ssh authentication set up.


Other Info
----------
When adding logos and other graphics it's best to minify/compress them first
before adding them to the repo. To do this, use imagemin.
It's available through npm and can be installed by running
```
sudo npm install -g imagemin
```
Then compress the image by running the command
```
imagemin my-bloated-image.img-ext web/assets/images/wherever-its-going
```
Once you do that, verify that it looks fine and then commit it to the repo.

