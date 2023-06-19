
from transformers import pipeline
import json

import time
start_time = time.time()
model_name = "SajjadAyoubi/bert-base-fa-qa"
qa_pipeline = pipeline("question-answering",
                       model=model_name, tokenizer=model_name)

text = "سلام من سجاد ایوبی هستم و به پردازش زبان طبیعی علاقه و سینما علاقه دارم"
questions = ["اسمم چیه؟", "علاقه مندیم چیه؟"]

results = []  # List to store the results

for question in questions:
    answer = qa_pipeline({"context": text, "question": question})
    results.append(answer)

# Save the results as JSON
with open("output.json", "w", encoding="utf-8") as file:
    json.dump(results, file, ensure_ascii=False)


# from transformers import pipeline
# import sys

# # Define a function to handle encoding errors
# def print_encoded(text):
#     try:
#         print(text)
#     except UnicodeEncodeError:
#         # If encoding error occurs, encode and decode with 'utf-8' to handle the characters
#         encoded_text = text.encode('utf-8', errors='replace').decode('utf-8')
#         print(encoded_text)

# model_name = "SajjadAyoubi/bert-base-fa-qa"
# qa_pipeline = pipeline("question-answering", model=model_name, tokenizer=model_name)

# text = "سلام من سجاد ایوبی هستم و به پردازش زبان طبیعی علاقه دارم"
# questions = ["اسمم چیه؟", "علاقه مندیم چیه؟"]

# for question in questions:
#     answer = qa_pipeline({"context": text, "question": question})
#     print_encoded(answer)

# from transformers import pipeline
# import sys

# # Define a function to handle encoding errors
# def print_encoded(text):
#     try:
#         print(text)
#     except UnicodeEncodeError:
#         # If encoding error occurs, encode and decode with 'utf-8' to handle the characters
#         encoded_text = text.encode('utf-8', errors='replace').decode('utf-8')
#         print(encoded_text)

# model_name = "SajjadAyoubi/bert-base-fa-qa"
# qa_pipeline = pipeline("question-answering", model=model_name, tokenizer=model_name)

# text = "سلام من سجاد ایوبی هستم و به پردازش زبان طبیعی علاقه دارم"
# questions = ["اسمم چیه؟", "علاقه مندیم چیه؟"]

# for question in questions:
#     answer = qa_pipeline({"context": text, "question": question})
#     for key, value in answer.items():
#         print_encoded(f"{key}: {value}")
