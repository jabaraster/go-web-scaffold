(function($) {
"use strict";

var Page = React.createClass({displayName: "Page",
    render: function() {
        return (
            React.createElement("div", {className: "Page"}, 
            "tex"
            )
        );
    }
});

React.render(
    React.createElement(Page, null),
    document.getElementById('page')
);

})(jQuery);
