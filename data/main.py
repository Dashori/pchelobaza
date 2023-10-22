import random
from faker import Faker
import os
fake = Faker()

USERS = 100
FARMS = 120
HONEY = 20
FARM_HONEY = 50
REQUESTS = 10
RECORDS = 100
CONF = 20
COMMENT = 20

file = open(str(RECORDS) + '.sql', 'w')

def date (year):
    month = random.randint(1, 12) 
    date = random.randint(1, 28)

    file.write(str(year) + "-"  + str(month) + "-" + str(date))

def dateTime (year):
    month = random.randint(1, 12) 
    date = random.randint(1, 28)
    time = random.randint(10, 20)

    file.write(str(year) + "-"  + str(month) + "-" + str(date) + " " + str(time) + ":00")

############################################################################
# user
file.write("insert into bee_user(login, password, name, surname, contact, registered_at, role) values \n")
roles = ["beeman", "beemaster"]

for i in range (0, USERS):
    file.write('(')
    
    # login
    file.write('\'')
    last_name = fake.last_name()
    file.write(last_name + str(i))
    file.write('\',')

    # password
    file.write('\'')
    file.write(str("12345"))
    file.write('\',')

    # name
    file.write('\'')
    file.write(fake.first_name())
    file.write('\',')

    # surname
    file.write('\'')
    file.write(last_name)
    file.write('\',')

    # email
    file.write('\'')
    file.write(fake.email())
    file.write('\',')

    # registered_at
    file.write('\'')
    year = random.randint(2018, 2020) 
    date(year)
    file.write('\',')

    # role
    if i < USERS/2 :
        file.write('\'')
        file.write(roles[0])
        file.write('\'')
    else:
        file.write('\'')
        file.write(roles[1])
        file.write('\'')

    if i != USERS - 1:
        file.write("),\n")
    else:
        file.write(");\n\n")

############################################################################
# farm
file.write("insert into bee_farm(id_user, name, description, address) values \n")

for i in range (0, FARMS):
    file.write('(')
    
    # id_user
    file.write('\'')
    file.write(str(random.randint(1, USERS)))
    file.write('\',')

    # name
    file.write('\'')
    farm = "farm " + str(fake.word() + str(i))
    file.write(farm)
    file.write('\',')

    # description
    file.write('\'')
    file.write("Welcome to my farm!")
    file.write('\',')

    # address
    file.write('\'')
    file.write(str(fake.country()) + ", " + str(fake.city()) + ", " + str(fake.street_name()) + ", " + str(fake.building_number()))
    file.write('\'')

    if i != FARMS - 1:
        file.write("),\n")
    else:
        file.write(");\n\n")

############################################################################
# farm_honey

file.write("insert into bee_farm_honey(id_farm, id_honey) values \n")

for i in range (0, FARM_HONEY):
    file.write('(')
    
    # id_farm
    file.write('\'')
    file.write(str(random.randint(1, FARMS)))
    file.write('\',')
    
    # id_honey
    file.write('\'')
    file.write(str(random.randint(1, HONEY)))
    file.write('\'')

    if i != FARM_HONEY - 1:
        file.write("),\n")
    else:
        file.write(");\n\n")

############################################################################
# bee_request

status = ["waiting", "approve", "rejected"]
file.write("insert into bee_request(id_user, description, status) values \n")

for i in range (0, REQUESTS):
    file.write('(')
    
    # id_user
    file.write('\'')
    file.write(str(int(i + USERS/2 + 1)))
    file.write('\',')

    # info
    file.write('\'')
    file.write("Please give me the beeman rules")
    file.write('\',')

    # status
    file.write('\'')
    file.write(str(status[random.randint(0, 2)]))
    file.write('\'')

    if i != REQUESTS - 1:
        file.write("),\n")
    else:
        file.write(");\n\n")

############################################################################
# bee_conf

file.write("insert into bee_conf(id_user, name, description, date, place, maximum_users, current_users) values \n")

for i in range (0, CONF):
    file.write('(')
    
    # id_user
    file.write('\'')
    file.write(str(int(i + USERS/2 + 1)))
    file.write('\',')

    #name
    file.write('\'')
    file.write("Conference" + str(i))
    file.write('\',')
    
    # description
    file.write('\'')
    file.write("Welcome to the conference!")
    file.write('\',')

    # date
    file.write('\'')
    year = random.randint(2021, 2023) 
    dateTime(year)
    file.write('\',')

    # place
    file.write('\'')
    file.write(str(fake.country()) + ", " + str(fake.city()) + ", " + str(fake.street_name()) + ", " + str(fake.building_number()))
    file.write('\',')
    
    # maximum_users
    file.write('\'')
    max_users = random.randint(10, 100) 
    file.write(str(max_users))
    file.write('\',')

    # current_users
    file.write('\'')
    # file.write(str(random.randint(0, max_users)))
    file.write(str(1))
    file.write('\'')

    if i != CONF - 1:
        file.write("),\n")
    else:
        file.write(");\n\n")


############################################################################
# user_conf

file.write("insert into bee_user_conf(id_user, id_conf) values \n")

for i in range (1, CONF):
    file.write('(')
    
    # id_user
    file.write('\'')
    file.write(str(i))
    file.write('\',')
    
    # id_conf
    file.write('\'')
    file.write(str(i))
    file.write('\'')

    if i != CONF - 1:
        file.write("),\n")
    else:
        file.write(");\n\n")

############################################################################
# bee_comment

file.write("insert into bee_comment(id_conf, id_user, time, description) values \n")

for i in range (1, COMMENT):
    file.write('(')
    
    # id_user
    file.write('\'')
    file.write(str(i))
    file.write('\',')
    
    # id_conf
    file.write('\'')
    file.write(str(i))
    file.write('\',')

    # time
    file.write('\'')
    dateTime(2023)
    file.write('\',')

    # description
    file.write('\'')
    file.write("It is my comment!")
    file.write('\'')


    if i != CONF - 1:
        file.write("),\n")
    else:
        file.write(");\n\n")