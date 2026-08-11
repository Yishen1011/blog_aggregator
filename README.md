# Pre-requisite
You will need to install **Go**, **Postgres** and **Goose** to run the program

Try running these commands

### **To install Go:**
#### Paste this into Terminal.app or a shell prompt, and press enter. Refer to [webi](https://webinstall.dev/golang/) for more information.
**Mac:**
curl -sS https://webi.sh/golang | sh; \
source ~/.config/envman/PATH.env

**Linux:**
curl -sS https://webi.sh/golang | sh; \
source ~/.config/envman/PATH.env

**Windows:**
curl.exe https://webi.ms/golang | powershell

### **To install Postgres:**
**macOS with brew**

*brew install postgresql@15*

**Linus / WSL (Debian)**

*sudo apt update*
*sudo apt install postgresql postgresql-contrib*

1.Run *psql --version* to make sure it's installed correctly.

2.(Linux / WSL only) Update postgres password: 

*sudo passwd postgres*

Enter a password, and be sure you won't forget it. You can just use something easy like postgres.

3.Start the Postgres server in the background

**Mac**

*brew services start postgresql@15*

**Linux** 

*sudo service postgresql start*

4.Connect to the server

Enter the psql shell:

**Mac**

*psql postgres*

**Linux**

*sudo -u postgres psql*

You should see a new prompt that looks like this:

`postgres=#`

### **To install Goose:**

go install github.com/pressly/goose/v3/cmd/goose@latest

Run *goose -version* to make sure it's installed correctly.

# To run the blog aggregator

### **Install gator module**

Run *go install .*

After installing module you can use these CLI commands.

*gator reset*
- Resets the users database

*gator register Name*
- Register the name into the users database

*gator login Name*
- Tries to login the name from the users database

*gator users*
- Lists all users in the users database

*gator agg (interval=optional)*
- Lists all posts from user feed each interval

*gator addfeed (title) (html website)*
- Add feeds to the current user

*gator feeds*
- List all feeds from the current user

*gator follow (website)*
- Follow website from the current user

*gator following*
- List all followed feed from the current user

*gator unfollow (website)*
- Unfollow website from the current user

*gator browse (limit=optional)*
- List all posts from RSS feed by the current user