from fastapi import FastAPI

app = FastAPI()

BOOKS = []


@app.get("/books")
async def read_all_books():
    return BOOKS

def main():
    print("Hello World")
    print("test")

if __name__ == "__main__":
    main()
