
$('input[type="checkbox"]#smart-fixed-header').click(function() {
	if($(this).is(':checked')) {
		//checked
		$.root_.addClass("fixed-header");
	}
	else {
		//unchecked
		$('input[type="checkbox"]#smart-fixed-ribbon')
			.prop('checked', false);
		$('input[type="checkbox"]#smart-fixed-navigation')
			.prop('checked', false);

		$.root_.removeClass("fixed-header");
		$.root_.removeClass("fixed-navigation");
		$.root_.removeClass("fixed-ribbon");

	}
});

/*
 * FIXED NAV
 */
$('input[type="checkbox"]#smart-fixed-navigation').click(function() {
	if($(this)
			.is(':checked')) {
		//checked
		$('input[type="checkbox"]#smart-fixed-header')
			.prop('checked', true);

		$.root_.addClass("fixed-header");
		$.root_.addClass("fixed-navigation");

		$('input[type="checkbox"]#smart-fixed-container')
			.prop('checked', false);
		$.root_.removeClass("container");

	}
	else {
		//unchecked
		$('input[type="checkbox"]#smart-fixed-ribbon')
			.prop('checked', false);
		$.root_.removeClass("fixed-navigation");
		$.root_.removeClass("fixed-ribbon");
	}
});

if(localStorage.getItem('sm-setmenu') == 'top') {
	$('#smart-topmenu')
		.prop('checked', true);
}
else {
	$('#smart-topmenu')
		.prop('checked', false);
}

/*
 * COLORBLIND FRIENDLY
 */

$('input[type="checkbox"]#colorblind-friendly')
	.click(function() {
		if($(this)
				.is(':checked')) {

			//checked
			$.root_.addClass("colorblind-friendly");

		}
		else {
			//unchecked
			$.root_.removeClass("colorblind-friendly");
		}
	});

/*
 * REFRESH WIDGET
 */
$("#reset-smart-widget").bind("click", function() {
	$('#refresh').click();
	return false;
});

