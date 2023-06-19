import json
import os

script_dir = os.path.dirname(os.path.abspath(__file__))
json_file = os.path.join(script_dir, "data.json")


with open(json_file) as f:
    data = json.load(f)


text = data["text"]
questions = data["questions"]

if 'questions' in data:
    del data['questions']

data["qa"] = [
    {
        "question": "balah balah",
        "answer": "answer"
    },
    {
        "question": "balah balah 2",
        "answer": "answer 2"
    }
]

with open(json_file, "w") as f:
    json.dump(data, f)

print("Python script execution complete.")
