#!/bin/sh

# Populate the hockeyeth configuration if it does not exist
for template in "${HKP_CONF_DIR}"/hockeypuck.conf.tmpl; do
  config="${HKP_CONF_DIR}/$(basename "$template" .tmpl)"
  [ ! -f "$config" ] &&
    cp "$template" "$config"
done

# Now call the normal entrypoint
/hockeypuck/bin/startup.sh
