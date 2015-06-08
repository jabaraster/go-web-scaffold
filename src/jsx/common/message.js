(function() {
"use strict";

var TransitionGroup = React.addons.CSSTransitionGroup;

window.Message = React.createClass({
    handleCloseClick: function() {
        this.props.onCloseClick();
    },
    render: function() {
        return (
            <div>
                {this.props.visible ?
                <div className="Message">
                    <div>
                        <span className="message-text">{this.props.text}</span>
                        <div className="tool-box">
                            <button onClick={this.handleCloseClick}><i className="glyphicon glyphicon-ok" /></button>
                        </div>
                    </div>
                </div>
                : null
                }
            </div>
        );
    }
});

})();
