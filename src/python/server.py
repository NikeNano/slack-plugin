import json
import os

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
            if 'hello' in args['template'].get('plugin', {}):
                small_message()
                self.reply({'node': {'phase': 'Succeeded', 'message': 'Hello template!'}})
            if 'test' in args['template'].get('plugin', {}):
                small_message()
                self.reply({'node': {'phase': 'Succeeded', 'message': 'Hello template!'}})
            else:
                self.reply({})
        else:
            self.unsupported()

# verifies if "the-welcome-channel" already exists
def small_message():
    token = os.environ["SLACK_BOT_TOKEN"]
    client = WebClient(token=token)

    # grab a list of all the channels in a workspace
    client.chat_postMessage(
        channel="test",
        text=f"Hello this is the bot{datetime.now()}"
    )

if __name__ == '__main__':
    httpd = HTTPServer(('', 4355), Plugin)
    httpd.serve_forever()