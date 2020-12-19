# book-show
Movie ticket buying Website

To use it on your local host you can use Following ways:

1) Using go in your pc locally -:

    - Download the zip file ,unzip it your desired folder.
    
    - Download and install go in pc.
    
    - now using "go get -v -u [name of repo]" ,download gin/gonic,bycrpt,cors,contrib static (use google to find the proper links to them).
    
    - now go to the backend folder through your cmd.
    
    - And run "go build server.go" this will build a server.exe file for you in the same folder.
    
    - in your cmd simply run "server.exe" this will run your backend service for website at localhost:3000 in your browser.
    
    - now go to the frontend folder through your cmd.
    
    - And run "go build server1.go" this will build a server1.exe file for you in the same folder.
    
    - in your cmd simply run "server1.exe" this will run your frontend service for website at localhost:2404 in your browser.
     
    - now install Xammpp for sql services.
    
    - start apache and mysql and from browser open phpmyadmin.
    
    - now make a new database in that namely showtime(you can change the name make sure u change the name in conn.go file too where we made database connection).
    
    - in showtime database import the file showtime.sql which you have in the the main folder of this project.
    
    - now finally u are good to go enjoy the website.(tho it has only client end only till now) :)
    
2) Using Docker desktop :-

    - easy as ever simply run the docker.yaml file which is the folder which contain this project ( docker-compose up ).
  
    - now you will have everything set up go to localhost 8080 for phpmyadmin .(username:root ,password:akshit).
  
    - now make a new database in that namely showtime(you can change the name make sure u change the name in conn.go file too where we made database connection).
  
    - in showtime database import the file showtime.sql which you have in the the main folder of this project.
    - you are good to go your front end will be on localhost:2404.
  
  Enjoy :)
