$(document).ready(function() {
		if($('-webkit-autofill').is()){
			$('input:-webkit-autofill').each(function(){
				$(this).siblings('label').addClass('used');
			});
		}
		

	  $('input:not(:hidden), select').on('keyup change blur click ', function(e) {
	    var $this = $(this);
	    if ($this.val() || $this.is(":focus") || (e.key == "Tab" || e.keyCode == 9)){
	    	$this.siblings('label').addClass('used')
		    $this.addClass('used');
	    }
	    else{
	    	$this.siblings('label').removeClass('used');
	    	$this.removeClass('used');
	    }
	  });

	  var $ripples = $('.ripples');

	  $ripples.on('click.Ripples', function(e) {

	    var $this = $(this);
	    var $offset = $this.parent().offset();
	    var $circle = $this.find('.ripplesCircle');

	    var x = e.pageX - $offset.left;
	    var y = e.pageY - $offset.top;

	    $circle.css({
	      top: y + 'px',
	      left: x + 'px'
	    });

	    $this.addClass('is-active');

	  });
	$('input, select').delay(1000).change();
	  $ripples.on('animationend webkitAnimationEnd mozAnimationEnd oanimationend MSAnimationEnd', function(e) {
	  	$(this).removeClass('is-active');
	  });

	});