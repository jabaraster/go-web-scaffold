(function($) {
"use strict";

window.ModalDialog = React.createClass({displayName: "ModalDialog",
    propTypes: {
        descriptor: React.PropTypes.string,
        visible: React.PropTypes.bool.isRequired,
        title: React.PropTypes.string.isRequired,
        cancelCaption: React.PropTypes.string,
        okCaption: React.PropTypes.string,
        onCancelClick: React.PropTypes.func.isRequired,
        onOkClick:  React.PropTypes.func.isRequired,
        onClosed: React.PropTypes.func,
        content: React.PropTypes.element.isRequired
    },
    getInitialState: function() {
        return {
            visible: false
        };
    },
    componentDidMount: function() {
        var d = this.props.descriptor ? '.' + this.props.descriptor : '.modal';
        $(document).on('hidden.bs.modal', '.ModalDialog ' + d, function(e) {
            if (this.props.onClosed) this.props.onClosed();
        }.bind(this));
    },
    componentDidUpdate: function(prevProps, prevState) {
        var d = this.props.descriptor ? '.' + this.props.descriptor : '.modal';
        if (this.props.visible) {
            $('.ModalDialog ' + d).modal('show');
        } else {
            $('.ModalDialog ' + d).modal('hide');
        }
    },
    handleCancelClick: function() {
        if (this.props.onCancelClick) this.props.onCancelClick();
    },
    handleOkClick: function() {
        this.props.onOkClick();
    },
    render: function() {
        return (
            React.createElement("div", {className: "ModalDialog"}, 
                React.createElement("div", {className: 'modal fade ' + (this.props.descriptor ? this.props.descriptor : ''), tabIndex: "-1"}, 
                    React.createElement("div", {className: "modal-dialog"}, 
                        React.createElement("div", {className: "modal-content"}, 
                            React.createElement("div", {className: "modal-header"}, 
                                React.createElement("button", {type: "button", className: "close", onClick: this.handleCancelClick}, 
                                    React.createElement("span", {"aria-hidden": "true"}, "×")
                                ), 
                                React.createElement("h4", {className: "modal-title"}, this.props.title)
                            ), 
                            React.createElement("div", {className: "modal-body"}, 
                                this.props.content
                            ), 
                            React.createElement("div", {className: "modal-footer"}, 
                                React.createElement("button", {type: "button", className: "btn btn-default", onClick: this.handleCancelClick}, 
                                    React.createElement("span", null, this.props.cancelCaption ? this.props.cancelCaption : 'キャンセル')
                                ), 
                                React.createElement("button", {type: "button", className: "btn btn-primary", onClick: this.handleOkClick}, 
                                    React.createElement("span", null, this.props.okCaption ? this.props.okCaption : 'OK')
                                )
                            )
                        )
                    )
                )
            )
        );
    }
});

})(jQuery);
