(function($) {
"use strict";

window.ModalDialog = React.createClass({
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
            <div className="ModalDialog">
                <div className={'modal fade ' + (this.props.descriptor ? this.props.descriptor : '')} tabIndex="-1">
                    <div className="modal-dialog">
                        <div className="modal-content">
                            <div className="modal-header">
                                <button type="button" className="close" onClick={this.handleCancelClick}>
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <h4 className="modal-title">{this.props.title}</h4>
                            </div>
                            <div className="modal-body">
                                {this.props.content}
                            </div>
                            <div className="modal-footer">
                                <button type="button" className="btn btn-default" onClick={this.handleCancelClick}>
                                    <span>{this.props.cancelCaption ? this.props.cancelCaption : 'キャンセル'}</span>
                                </button>
                                <button type="button" className="btn btn-primary" onClick={this.handleOkClick}>
                                    <span>{this.props.okCaption ? this.props.okCaption : 'OK'}</span>
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
});

})(jQuery);
