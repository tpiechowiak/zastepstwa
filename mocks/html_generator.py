import names
from jinja2 import Template

from random import choice, randint

GROUPS = ["3E", "1C|N1", "1GB", "2Z", "1A", "4D", "2A|1/2"]
LESSONS = ["Niemiecki", "Programowanie", "Fajny przedmiot", "Zażółć"]
ROOMS = [str(randint(20, 40)) for i in range(20)] + ["SG1", "@"]
TEACHERS = [names.get_full_name() for i in range(20)]


def render(subs):
    with open("template.html") as f:
        template = Template(f.read())

    return template.render(subs=subs)


def generate_subs(n=10):
    return [
        {
            "lesson_number": randint(0, 7),
            "regular_teacher": choice(TEACHERS),
            "group": choice(GROUPS),
            "lesson_name": choice(LESSONS),
            "room": choice(ROOMS),
            "substitute_teacher": choice(TEACHERS + ["Uczniowie przychodzą później", "Uczniowie zwolnieni do domu"])
        }
        for i in range(n)
    ]


if __name__ == "__main__":
    subs = generate_subs()
    print(render(subs))
