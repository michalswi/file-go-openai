### **analyze_html** 

pattern is available [here](./pattern) .  

To limit/narrow down search:
```
sed -n '/<body/,/<\/body>/p' /tmp/index.html > /tmp/index.html_body
sed -n '/<head/,/<\/head>/p' /tmp/index.html > /tmp/index.html_head
sed -n '/<script/,/<\/script>/p' /tmp/index.html > /tmp/index.html_script
```
