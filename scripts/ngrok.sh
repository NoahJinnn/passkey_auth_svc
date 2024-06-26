#!/bin/sh
PROFILE=$1
echo "⚡️ Kill old ngrok"
ngrok_pid=$(pgrep ngrok)
if [ -n "$ngrok_pid" ]; then
  kill -9 "$ngrok_pid"
  echo "ngrok process killed with PID: $ngrok_pid"
else
  echo "ngrok process not found"
fi
sleep 1

# Start NGROK in background
echo "⚡️ Starting ngrok"
ngrok http 17000 >/dev/null &
sleep 1

NGROK_REMOTE_URL="$(curl http://localhost:4040/api/tunnels | jq ".tunnels[0].public_url")" # get ngrok url from ngrok's tunnel response



# Get NGROK dynamic URL from its own exposed local API
if test -z "${NGROK_REMOTE_URL}"
then
  echo "❌ ERROR: ngrok doesn't seem to return a valid URL (${NGROK_REMOTE_URL})."
  exit 1
else
  if [ "$NGROK_REMOTE_URL" == "null" ]; then
    echo "❌ ERROR: ngrok doesn't seem to return a valid URL (${NGROK_REMOTE_URL})."
    exit 1
  fi
fi


# Trim double quotes from variable
NGROK_REMOTE_URL=$(echo ${NGROK_REMOTE_URL} | tr -d '"')
# If http protocol is returned, replace by https
NGROK_REMOTE_URL=${NGROK_REMOTE_URL/http:\/\//https:\/\/}
# Trim https protocol from variable
NGROK_REMOTE_URL=${NGROK_REMOTE_URL/https:\/\//}

echo "\n🌍 Your ngrok remote URL is 👉 ${bold}${NGROK_REMOTE_URL}"

doppler configure set config=$PROFILE project=passkey_auth_svc
doppler secrets set AUTH_DOMAIN $NGROK_REMOTE_URL

doppler configure set config=$PROFILE project=passkey-auth-service
doppler secrets set AUTH_DOMAIN $NGROK_REMOTE_URL