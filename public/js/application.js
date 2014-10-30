window.addEventListener('load', function() {
  "use strict"

  var button = document.querySelector('button');
  button.addEventListener('click', function() {
    var togglable = document.querySelectorAll('.js-togglable');

    [].forEach.call(togglable, function(elem) {
      elem.classList.toggle('hidden');
    });
  });
});