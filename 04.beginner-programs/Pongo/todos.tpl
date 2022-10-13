{% for todo in todos -%}
    {% if todo.Done %}
        {{- todo.Title -}}
    {% endif %}
{% endfor %}
