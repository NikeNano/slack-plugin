import json
import os
import ast

from slack_sdk import WebClient
from datetime import datetime
from http.server import BaseHTTPRequestHandler, HTTPServer
from pprint import pprint

class Plugin(BaseHTTPRequestHandler):

    def args(self):
        return json.loads(self.rfile.read(int(self.headers.get('Content-Length'))))

    def reply(self, reply):
        self.send_response(200)
        self.end_headers()
        self.wfile.write(json.dumps(reply).encode("UTF-8"))

    def unsupported(self):
        self.send_response(404)
        self.end_headers()

    def do_POST(self):
        if self.path == '/api/v1/template.execute':
            args = self.args()
            pprint(args)
            small_message(args)
            if 'hello' in args['template'].get('plugin', {}):
                self.reply({'node': {'phase': 'Succeeded', 'message': 'Hello template!'}})
            else:
                self.reply({})
        else:
            self.unsupported()

# verifies if "the-welcome-channel" already exists
def small_message(payload: dict):
    token = os.environ["SLACK_BOT_TOKEN"]
    client = WebClient(token=token)

    vals = ast.literal_eval(payload["template"]["plugin"]["hello"])
    if "channel" not in vals:
        return 
    client.chat_postMessage(
        channel=vals.get("channel", ""),
        text=f"Hello this is the bot: {datetime.now()}, message: {vals.get('text', '')}"
    )

if __name__ == '__main__':
    httpd = HTTPServer(('', 4355), Plugin)
    httpd.serve_forever()