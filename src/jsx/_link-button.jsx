var React = require('react');
var Bootstrap = require('react-bootstrap');
var Button = Bootstrap.Button;
var Glyphicon = Bootstrap.Glyphicon;

var LinkButton = React.createClass({
    propTypes: {
        href: React.PropTypes.string.isRequired,
        text: React.PropTypes.string,
        glyph: React.PropTypes.string,
        style: React.PropTypes.string
    },
    handleClick: function() {
        location.href = this.props.href;
    },
    render: function() {
        var glyph = this.props.glyph ? (<Glyphicon glyph={this.props.glyph} />) : (<span/>);
        return (
            <Button bsStyle={this.props.style} onClick={this.handleClick}>
                {glyph}
                {this.props.text}
            </Button>
        );
    }
});

module.exports = LinkButton;
