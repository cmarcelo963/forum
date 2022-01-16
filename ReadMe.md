Forum

Steps:
1. Create Server
    a. Homepage Template 
        i. Find UI Design
            aa. Logo Top-Left
            bb. Search Bar Middle Top
            cc. Timer Session - User Login Top-Right
            dd. 

2. Authentication
    a. Front-end
        i. Register new user
        ii. Log-in existing user
    b. Back-end
        i. Create the schema
        ii. Create the database
        iii. Register: Handle post request from front-end
            aa. Parse the form
            bb. Validate data
            cc. Connect with databse - check user doesn't exist
            dd. Create user
            ee. Send confirmation to front-end
        iv. Log-in: Handle post request from front-end
            aa. Parse the form
            bb. Validate data
            cc. Connect with database - match user if password
            dd. Send template with active session - cookies