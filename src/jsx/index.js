(function($) {
"use strict";

var Page = React.createClass({
    render: function() {
        return (
            <div className="Page">
            </div>
        );
    }
});

React.render(
    <Page />,
    document.getElementById('page')
);

})(jQuery);
