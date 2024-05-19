data "external" "env_vars" { 
    program = ["../scripts/set_envs_tf.sh"]
}
