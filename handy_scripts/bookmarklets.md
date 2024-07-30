1. **[INPUT-FIELDS] Enter Go-To Payload and Reload**
2. **[INPUT-FIELDS] Enter HTML-only Payload and Reload**
3. **[URL-PARAMS] Set Params to Go-To Payload and Reload**
4. **[URL-PARAMS] Set Params to HTML-only Payload and Reload**
5. **[COPY] SSTI Broken Syntax Payload**
6. **[COPY] BXSS Payload 1**
7. **[COPY] BXSS Payload 2**
8. **[COPY] BXSS Payload 3**
9. **[COPY] XSS Polyglot**
10. **[COPY] Go-To Payload**

Here’s how you can create bookmarklets for each case:

### 1. [INPUT-FIELDS] Enter Go-To Payload and Reload

```javascript
javascript:(function(){
  var payload = 'YOUR_GO_TO_PAYLOAD_HERE';
  document.querySelectorAll('input, textarea').forEach(function(input) {
    input.value = payload;
  });
  location.reload();
})();
```

### 2. [INPUT-FIELDS] Enter HTML-only Payload and Reload

```javascript
javascript:(function(){
  var payload = 'YOUR_HTML_ONLY_PAYLOAD_HERE';
  document.querySelectorAll('input, textarea').forEach(function(input) {
    input.value = payload;
  });
  location.reload();
})();
```

### 3. [URL-PARAMS] Set Params to Go-To Payload and Reload

```javascript
javascript:(function(){
  var payload = 'YOUR_GO_TO_PAYLOAD_HERE';
  var params = new URLSearchParams(window.location.search);
  params.forEach((value, key) => {
    params.set(key, payload);
  });
  window.location.search = params.toString();
})();
```

### 4. [URL-PARAMS] Set Params to HTML-only Payload and Reload

```javascript
javascript:(function(){
  var payload = 'YOUR_HTML_ONLY_PAYLOAD_HERE';
  var params = new URLSearchParams(window.location.search);
  params.forEach((value, key) => {
    params.set(key, payload);
  });
  window.location.search = params.toString();
})();
```

### 5. [COPY] SSTI Broken Syntax Payload

```javascript
javascript:(function(){
  var payload = 'YOUR_SSTI_BROKEN_SYNTAX_PAYLOAD_HERE';
  navigator.clipboard.writeText(payload).then(function() {
    alert('SSTI Broken Syntax Payload copied to clipboard');
  }, function(err) {
    alert('Could not copy payload: ', err);
  });
})();
```

### 6. [COPY] BXSS Payload 1

```javascript
javascript:(function(){
  var payload = 'YOUR_BXSS_PAYLOAD_1_HERE';
  navigator.clipboard.writeText(payload).then(function() {
    alert('BXSS Payload 1 copied to clipboard');
  }, function(err) {
    alert('Could not copy payload: ', err);
  });
})();
```

### 7. [COPY] BXSS Payload 2

```javascript
javascript:(function(){
  var payload = 'YOUR_BXSS_PAYLOAD_2_HERE';
  navigator.clipboard.writeText(payload).then(function() {
    alert('BXSS Payload 2 copied to clipboard');
  }, function(err) {
    alert('Could not copy payload: ', err);
  });
})();
```

### 8. [COPY] BXSS Payload 3

```javascript
javascript:(function(){
  var payload = 'YOUR_BXSS_PAYLOAD_3_HERE';
  navigator.clipboard.writeText(payload).then(function() {
    alert('BXSS Payload 3 copied to clipboard');
  }, function(err) {
    alert('Could not copy payload: ', err);
  });
})();
```

### 9. [COPY] XSS Polyglot

```javascript
javascript:(function(){
  var payload = 'YOUR_XSS_POLYGLOT_HERE';
  navigator.clipboard.writeText(payload).then(function() {
    alert('XSS Polyglot copied to clipboard');
  }, function(err) {
    alert('Could not copy payload: ', err);
  });
})();
```

### 10. [COPY] Go-To Payload

```javascript
javascript:(function(){
  var payload = 'YOUR_GO_TO_PAYLOAD_HERE';
  navigator.clipboard.writeText(payload).then(function() {
    alert('Go-To Payload copied to clipboard');
  }, function(err) {
    alert('Could not copy payload: ', err);
  });
})();
```

Replace `YOUR_GO_TO_PAYLOAD_HERE`
