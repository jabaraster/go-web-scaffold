(function() {
"use strict";

window.Indicator = React.createClass({displayName: "Indicator",
    componentDidUpdate: function() {
        var indicator = $(React.findDOMNode(this.refs.indicator));
        if (!this.props.active) {
            indicator.activity(false).hide();
            return;
        }

        var button = $(this.props.buttonRef);
        indicator.height(button.outerHeight());
        indicator.width(button.outerWidth());
        indicator.css(button.position());
        indicator.show().activity();
    },
    handle: function(e) {
        e.stopPropagation();
    },
    render: function() {
        var style = {
            position: "absolute",
            zIndex: "900",
            cursor: "default",
            background: "rgba(255,255,255,.3)"
        };
        return React.createElement("span", {className: "Indicator", ref: "indicator", onClick: this.handle, style: style})
    },
});

})();
