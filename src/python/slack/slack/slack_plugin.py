import os
from slack_sdk import WebClient
from datetime import datetime


# verifies if "the-welcome-channel" already exists
def channel_exists():
    token = os.environ["SLACK_BOT_TOKEN"]
    client = WebClient(token=token)

    # grab a list of all the channels in a workspace
    client.chat_postMessage(
        channel="test",
        text=f"Hello this is the bot{datetime.now()}"
    )

if "__main__" == __name__:
    channel_exists()