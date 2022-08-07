import pandas as pd
from sqlalchemy import create_engine
from os import getenv

host = getenv('DBHOST')
user = getenv('DBUSER')
passwd = getenv('DBPASS')
port = getenv('DBPORT')
database = getenv('DBNAME')


engine = create_engine(
        url=f'postgresql://{user}:{passwd}@{host}:{port}/{database}',
        echo=False)

df = pd.read_csv('data.csv')

df.to_sql(name='dataset', con=engine, if_exists='replace', index=False)
