# from transformers import MT5ForConditionalGeneration, MT5Tokenizer
# model_size = "base"
# model_name = f"persiannlp/mt5-{model_size}-parsinlu-squad-reading-comprehension"
# tokenizer = MT5Tokenizer.from_pretrained(model_name)
# model = MT5ForConditionalGeneration.from_pretrained(model_name)

# -*- coding: utf-8 -*-

# from transformers import T5Tokenizer, TFMT5ForConditionalGeneration

# model_name = "google/mt5-base"
# tokenizer = T5Tokenizer.from_pretrained(model_name)
# model = TFMT5ForConditionalGeneration.from_pretrained(model_name)


# def run_model(paragraph, question, **generator_args):
#     input_ids = tokenizer.encode(str(question) + " \n " + str(paragraph) , return_tensors="pt")
#     res = model.generate(input_ids, **generator_args)
#     output = tokenizer.batch_decode(res, skip_special_tokens=True)
#     print(output)
#     return output
# run_model(
# "سلام من پویا هستم ، مهندسی کامپیوتر میخونم در دانشگاه شیراز و از رشتم راضی هستم ولی از دانشگاهم نه  ؟".encode("utf-8"),
# "دانشگاه شیراز ".encode("utf-8")
# )

# run_model(
#     "یک شی را دارای تقارن می‌نامیم زمانی که ان شی را بتوان به دو یا چند قسمت تقسیم کرد که آن‌ها قسمتی از یک طرح سازمان یافته باشند یعنی بر روی شکل تنها جابجایی و چرخش و بازتاب و تجانس انجام شود و در اصل شکل تغییری به وجود نیایید آنگاه ان را تقارن می‌نامیم مرکز تقارن:اگر در یک شکل نقطه‌ای مانندA وجود داشته باشد که هر نقطهٔ روی شکل (محیط) نسبت به نقطه یAمتقارن یک نقطهٔ دیگر شکل (محیط) باشد، نقطهٔ Aمرکز تقارن است. یعنی هر نقطه روی شکل باید متقارنی داشته باشد شکل‌های که منتظم هستند و زوج ضلع دارند دارای مرکز تقارند ولی شکل‌های فرد ضلعی منتظم مرکز تقارن ندارند. متوازی‌الأضلاع و دایره یک مرکز تقارن دارند ممکن است یک شکل خط تقارن نداشته باشد ولی مرکز تقارن داشته باشد. (منبع:س. گ)".encode("utf-8"),
#     "اشکالی که یک مرکز تقارن دارند"
# )
# import json
# from transformers import T5Tokenizer, TFMT5ForConditionalGeneration

# model_name = "google/mt5-base"
# tokenizer = T5Tokenizer.from_pretrained(model_name)
# model = TFMT5ForConditionalGeneration.from_pretrained(model_name)


# def run_model(paragraph, question, **generator_args):
#     input_ids = tokenizer.encode(str(question) + " \n " + str(paragraph), return_tensors="pt")
#     res = model.generate(input_ids, **generator_args)
#     output = tokenizer.batch_decode(res, skip_special_tokens=True)
#     return output


# paragraph_text = "سلام اسم من پویا است"
# question_text = "اسم من چی هست ؟"

# output = run_model(
#     paragraph_text,
#     question_text
# )

# # Save the output to a JSON file
# output_data = {
#     "paragraph": paragraph_text,
#     "question": question_text,
#     "output": output
# }

# with open("output.json", "w", encoding="utf-8") as f:
#     json.dump(output_data, f, ensure_ascii=False, indent=4)


import json
from transformers import MT5ForConditionalGeneration, MT5Tokenizer
model_size = "base"
model_name = f"persiannlp/mt5-{model_size}-parsinlu-squad-reading-comprehension"
tokenizer = MT5Tokenizer.from_pretrained(model_name)
model = MT5ForConditionalGeneration.from_pretrained(model_name)


def run_model(paragraph, question, **generator_args):
    input_ids = tokenizer.encode(question + "\n" + paragraph, return_tensors="pt")
    res = model.generate(input_ids, **generator_args)
    output = tokenizer.batch_decode(res, skip_special_tokens=True)
    return output

paragraph_text = "بیشتر خفاش‌ها شب‌زی‌اند. آن‌ها در طول روز یا خوابند یا به پاکسازی بدن خود می‌پردازند و در هنگام شب به شکار می‌روند. ابزار مسیریابی و شکار خفاش‌ها در تاریکی تا دههٔ ۱۷۹۰ کاملاً ناشناخته بود تا اینکه کشیش و زیست شناس ایتالیایی لازارو اسپالانزانی به مجموعه آزمایش‌هایی بر روی خفاش‌های کور دست زد. این خفاش‌ها در یک اتاق کاملاً تاریک گذاشت و مسیر آن‌ها را با نخ‌های ابریشمی پُر پیچ و خم کرد. حتی در تاریکی مطلق هم شبکورها راه خود را در آن مسیر پر پیچ و خم پیدا کرده بودند به همین دلیل او نتیجه گرفت که ابزار راه‌یابی شبکورها چیزی غیر از چشمانشان است."
question_text = "چرا خفاش در شب بیدار است؟"

output = run_model(paragraph_text, question_text)

# Save the output to a JSON file
output_data = {
    "paragraph": paragraph_text,
    "question": question_text,
    "output": output
}

with open("output.json", "w", encoding="utf-8") as f:
    json.dump(output_data, f, ensure_ascii=False, indent=4)
