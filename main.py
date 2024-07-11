from jinja2 import Environment, FileSystemLoader

# Specify the directory containing the templates
file_loader = FileSystemLoader('templates')

# Create an environment
env = Environment(loader=file_loader)

# Load the template
template = env.get_template('index.jinja2.html')

# Render the template with a context
rendered_template = template.render(name="MyName", items=["Item1", "Item2", "Item3"])

print(rendered_template)