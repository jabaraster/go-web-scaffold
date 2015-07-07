'use strict';

var React = require('react/addons');

var Message = React.createClass({
    propTypes: {
        messages: React.PropTypes.array.isRequired
    },
    render: function() {
        var rows = this.props.messages.map((msg, idx) => {
            return (
                <li key={`Message_${idx}`}>{msg}</li>
            );
        });
        var classes = React.addons.classSet({
            Message: true,
            hidden: this.props.messages.length === 0
        });
        return (
            <div className={classes}>
                <ul>
                    {rows}
                </ul>
            </div>
        );
    }
});

module.exports = Message;
