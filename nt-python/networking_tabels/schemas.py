from pydantic import BaseModel


class User(BaseModel):
    username: str
    uuid: str


class Table(BaseModel):
    uuid: str
    attendees: list[User]
    capacity: int


class TableArrangement(BaseModel):
    iteration: Table
    table_arrangement: list[Table]


class Session(BaseModel):
    uuid: str
    leader: User
    attendees: list[User]
    tables: list[Table]
    name: str
    description: str
    create_date: str
