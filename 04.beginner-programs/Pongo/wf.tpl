{% for word in words -%}
    {{ word }} has {{ word | length }} characters
{% endfor %}
