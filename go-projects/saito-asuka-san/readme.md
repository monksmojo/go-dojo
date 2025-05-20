<h1 align='center'>Saito Asuka San</h1>

<h3 align='center'>Expense Management System</h3>


#### UML  Use Case Diagram


#### Journey

**User Journey**
-  User logs in the system
-  User picks the date to record the expense
-  User adds the expense type
-  User adds the expense cycle
-  User adds the expense description
-  User adds the expense price  



#### System Terminology

1. **Expense**
   - **What is an expense**
      - An expense is the transaction which include deduction of money from your wallet or bank account for a expense type



#### System Unique Selling Point

1. **System Features/USP**
     - System has following default types
        - GROCERIES
        - FOOD
        - TRANSPORT
        - AMENITIES
        - MISC
        - RENT
    - System can support custom type
    - System has expense cycle
      - one time:-
        - expense transaction type done is unique and will not reoccur next month
      - reoccurring:-
        - expense transaction type done is common and reoccur every month
        - default expense cycle
        - reoccurring expense cycle help us to create expense forecast
    - Summary of expense of previous, current and next month for better expense planning and tracking
    - grouped summaries based on types     
    - expense forecasting    


#### DataBase Design

1. **user schema**
    - id varchar(255) primary key
    - name varchar(255) not null
    - email varchar(255) not null unique
    - password varchar(255) not null (initially auto generate)
    - expense_table_name varchar(255) not null unique (email_handle + "expense_table"+time_stamp_in_long)
    - country varchar(255) not null
    - currency varchar(255) not null

2. **user_type schema**
    - id string varchar(255) primary key
    - name string varchar(255) unique case insensitive 
    - user_id foreign_key references user(id) 

3. **type schema**
    - id string varchar(255) primary key
    - name string varchar(255) unique case insensitive 

4. **Expense schema**
   - id varchar(255) primary_key
   - timestamp long not null
   - day int not null
   - month int not null
   - year int not null
   - user_type_id varchar(255) foreign key reference user_type(id)
   - expense cycle varchar(255) default one-time not null
   - description varchar(755) not null
   - price double not null
   - record_inserted_on timestamp  


#### tech stack

1. BE (Back End)
   1. Go
   2. Gin
   3. Gorm
2. FE (Front End)
   1. HTMX
3. DATABASE
   1. MYSQL


