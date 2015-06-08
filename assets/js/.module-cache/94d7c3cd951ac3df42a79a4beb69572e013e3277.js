(function() {
"use strict";

var TransitionGroup = React.addons.CSSTransitionGroup;

window.Message = React.createClass({displayName: "Message",
    handleCloseClick: function() {
        this.props.onCloseClick();
    },
    render: function() {
        return (
            React.createElement("div", null, 
                this.props.visible ?
                React.createElement("div", {className: "Message"}, 
                    React.createElement("div", null, 
                        React.createElement("span", {className: "message-text"}, this.props.text), 
                        React.createElement("div", {className: "tool-box"}, 
                            React.createElement("button", {onClick: this.handleCloseClick}, React.createElement("i", {className: "glyphicon glyphicon-ok"}))
                        )
                    )
                )
                : null
                
            )
        );
    }
});

})();
