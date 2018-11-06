$(document).ready(function(){
	$('#everything-right').on('change', '#group-field', function(){
		
		var value = $('option[value="'+$(this).val()+'"]').text();
		if(value == "Other"){
			$('#other-group').fadeIn().find('input').focus();
		}else{
			$('#other-group').fadeOut();
		}
	});

	$('#register-form').submit(function(e){
		e.preventDefault();
		console.log('submit');
		var value = $('option[value="'+$('#group-field').val()+'"]').text();
		
		if(value == "Other"){
			var groupVal = $('input[name="OtherName"').val();
			if(groupVal){
				$('#register-form').off('submit').submit();
				return;
			}

			$('#other-group').focus().find('.error').fadeIn();
		}else{
			$('#other-group').find('.error').fadeOut();
			$('#register-form').off('submit').submit();
		}
	});
	$('#group-field').change();
});